import 'dart:js';

import 'package:bot_toast/bot_toast.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart' show CallOptions, ResponseFuture;
import 'package:protobuf/protobuf.dart' show GeneratedMessage;

import '../l10n.dart';
import '../pgde/error/error.pb.dart';
import '../validate.dart';

import 'data.dart';

typedef ByFieldNumberFunc = String Function(int fieldNumber);
typedef NoNeedAskCallback = bool Function();

abstract class ErrorCoder {
  String $byErrorCode(int number);
}

class ErrorText extends StatelessWidget {
  final int fieldNumber;
  final ByFieldNumberFunc byNumber;
  final int code;
  final ErrorCoder coder;
  final String message;

  const ErrorText(
      this.fieldNumber, this.byNumber, this.code, this.coder, this.message,
      {Key key})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return fieldNumber == 0
        ? Container()
        : Text('${coder.$byErrorCode(code)}' +
            (message.isEmpty ? '' : ': $message'));
  }
}

class GrpcErrorRetry extends StatelessWidget {
  final Object err;
  final VoidCallback onRetry;

  const GrpcErrorRetry({Key key, @required this.err, @required this.onRetry})
      : super(key: key);

  @override
  Widget build(BuildContext context) => Center(
        child: Column(
          children: <Widget>[
            Text(
              // TODO add l10n?
              '$err',
              style: const TextStyle(color: Colors.red),
            ),
            RaisedButton(
                onPressed: onRetry,
                child: Text(
                  PgdeLocalizations.of(context).formRetry,
                )),
          ],
        ),
      );
}

Future<bool> onStayOrWillPop(BuildContext context, VoidCallback onGoBack,
    {bool noNeedAsk}) {
  if (noNeedAsk == true) {
    Navigator.of(context).pop();
    return Future.value(false);
  }

  final l10n = PgdeLocalizations.of(context);
  showDialog<void>(
    context: context,
    barrierDismissible: true,
    builder: (BuildContext dialogContext) {
      return AlertDialog(
        title: Text(l10n.formBackAlertTitle),
        content: SingleChildScrollView(
          child: ListBody(
            children: <Widget>[
              Text(l10n.formBackAlertContent),
            ],
          ),
        ),
        actions: <Widget>[
          FlatButton(
            child: Text(l10n.formBackAlertStay),
            onPressed: () => Navigator.of(dialogContext).pop(),
          ),
          FlatButton(
            child: Text(l10n.formBackAlertGoBack),
            onPressed: () {
              Navigator.of(dialogContext).pop();
              onGoBack();
            },
          ),
        ],
      );
    },
  );
  return Future.value(false);
}

typedef SaveFunc<T extends GeneratedMessage> = ResponseFuture<BackendError>
    Function(T request, {CallOptions options});

class FormPage<E> extends StatefulWidget {
  final bool dev;
  final String title;
  final SaveFunc save;
  final ByFieldNumberFunc byNumber;
  final ErrorCoder coder;
  final Inputs inputs;

  const FormPage(
      {Key key,
      this.dev = false,
      this.title,
      this.save,
      this.byNumber,
      this.coder,
      this.inputs})
      : super(key: key);

  @override
  FormPageState createState() => FormPageState();
}

class FormPageState extends State<FormPage> {
  String _route;
  String route(BuildContext context) =>
      _route ??= ModalRoute.of(context).settings.name;

  Future<BackendError> _saving;

  void _doSave(BuildContext context) {
    setState(() {
      _saving = widget.save(widget.inputs.data.editing).then((err) {
        if (err == null) {
          Navigator.of(context).pop();
        }
        return err;
      });
    });
  }

  @override
  void initState() {
    super.initState();
  }

  @override
  void dispose() {
    _route = null;
    _saving = null;
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title + (widget.dev ? '-${route(context)}' : '')),
        leading: BackButton(),
      ),
      body: Container(
        margin: const EdgeInsets.all(15.0),
        child: FutureBuilder(
          future: _saving,
          builder:
              (BuildContext context, AsyncSnapshot<BackendError> snapshot) {
            switch (snapshot.connectionState) {
              case ConnectionState.waiting:
              case ConnectionState.active:
                return const Center(child: CircularProgressIndicator());
              case ConnectionState.none:
              case ConnectionState.done:
                if (snapshot.hasError)
                  return GrpcErrorRetry(
                    err: snapshot.error,
                    onRetry: () => _doSave(context),
                  );

                final err = snapshot.data;
                return Column(
                  children: <Widget>[
                    StatefulForm(
                      inputs: widget.inputs,
                      onSave: () => _doSave(context),
                    ),
                    err == null
                        ? Container()
                        : ErrorText(
                            err.fieldNumber,
                            widget.byNumber,
                            err.code,
                            widget.coder,
                            err.message,
                          ),
                  ],
                );
            }
            assert(
                false, 'unknown ConnectionState: ${snapshot.connectionState}');
            return null;
          },
        ),
      ),
    );
  }
}

class StatefulForm<E> extends StatefulWidget {
  final Inputs inputs;
  final VoidCallback onSave;

  const StatefulForm({Key key, this.inputs, this.onSave}) : super(key: key);

  @override
  State<StatefulWidget> createState() => StatefulFormState();
}

class StatefulFormState extends State<StatefulForm> {
  final GlobalKey<FormState> _fbKey = GlobalKey<FormState>();

  GeneratedMessage _draft;

  GeneratedMessage get draft => _draft ??=
      widget.inputs.data.editing?.clone() ?? widget.inputs.createEmptyDraft();

  bool _needNotSave() => _draft == widget.inputs.data.saved;

  void _doSave(BuildContext context, PgdeLocalizations l10n) {
    if (!_fbKey.currentState.validate()) {
      BotToast.showText(text: l10n.formHasError);
      return;
    }
    if (_needNotSave()) {
      Navigator.of(context).pop();
      return;
    }
    if (widget.inputs.data.isRoot) return setState(widget.onSave);
    widget.inputs.data.saveSelf(draft);
    Navigator.of(context).pop();
  }

  @override
  Widget build(BuildContext context) {
    final ml10n = MaterialLocalizations.of(context);
    final l10n = PgdeLocalizations.of(context);
    final theme = Theme.of(context);
    return Column(
      children: <Widget>[
        Form(
          key: _fbKey,
          // autovalidate: true,
          // readonly: true,
          onWillPop: () => onStayOrWillPop(
            context,
            () => Navigator.of(context).pop(),
            noNeedAsk: _needNotSave(),
          ),
          child: Column(children: widget.inputs.build(context)),
        ),
        Row(
          children: <Widget>[
            Expanded(
              child: MaterialButton(
                color: theme.accentColor,
                child: Text(
                  ml10n.okButtonLabel,
                  style: TextStyle(color: Colors.white),
                ),
                onPressed: () => _doSave(context, l10n),
              ),
            ),
            SizedBox(width: 20),
            Expanded(
              child: MaterialButton(
                color: theme.accentColor,
                child: Text(
                  l10n.formReset,
                  style: TextStyle(color: Colors.white),
                ),
                onPressed: () => _fbKey.currentState.reset(),
              ),
            ),
          ],
        ),
      ],
    );
  }
}

abstract class InOneofInputs {
  List<Widget> build(BuildContext context);
}

class InputsWidget<T extends GeneratedMessage> extends StatelessWidget {
  final InOneofInputs inputs;
  const InputsWidget(this.inputs, {Key key}) : super(key: key);
  @override
  Widget build(BuildContext context) => Column(children: inputs.build(context));
}

typedef CreateEmptyDraftFunc<T extends GeneratedMessage> = T Function();

// also act as EmbedMessageOneofInputs by implementing InOneofInputs
abstract class Inputs<T extends GeneratedMessage> implements InOneofInputs {
  static FormFieldValidator catchError(VoidCallback fn) {
    return (v) {
      try {
        fn();
      } catch (e) {
        return e.toString();
      }
      return null;
    };
  }

  final FormMessageData data;
  Inputs(this.data);

  T _draft;
  T get draft => _draft ??= data.editing?.clone() ?? createEmptyDraft();

  bool _needNotSave() => _draft == data.saved;

  InputDecoration decoration(BuildContext context) =>
      InputDecoration().applyDefaults(Theme.of(context).inputDecorationTheme);

  ValidateInfo<T> vi(BuildContext context) => ValidateInfo<T>(
      MaterialLocalizations.of(context), PgdeLocalizations.of(context), draft);

  CreateEmptyDraftFunc<T> get createEmptyDraft;

  @override
  List<Widget> build(BuildContext context);
}

typedef FieldInOneofBuilder = Widget Function(FormFieldValidator validator);

typedef CreateValidatorFunc<T extends GeneratedMessage> = GeneratedValidator<T>
    Function(BuildContext context, ValidateInfo<T> vi);

typedef CreateDecorationFunc = InputDecoration Function(
    BuildContext context, String errorText);

// non-embed oneof inputs
class MessageFieldInOneofInputs<T extends GeneratedMessage>
    implements InOneofInputs {
  // data from childMessage
  final FormMessageData data;
  final CreateEmptyDraftFunc<T> createEmptyDraft;
  final CreateValidatorFunc<T> createValidator;
  final CreateDecorationFunc createDecoration;

  MessageFieldInOneofInputs(
      {@required this.data,
      @required this.createEmptyDraft,
      @required this.createValidator,
      @required this.createDecoration});

  @override
  List<Widget> build(BuildContext context) {
    final draft = data.editing?.clone() ?? createEmptyDraft();
    final l10n = PgdeLocalizations.of(context);
    final vi = ValidateInfo(MaterialLocalizations.of(context), l10n, draft);
    String errorText;
    try {
      createValidator(context, vi).assertProto();
    } catch (e) {
      errorText = e.toString();
    }
    return [
      InputDecorator(
        decoration: createDecoration(context, errorText)
            .applyDefaults(Theme.of(context).inputDecorationTheme),
        child: MaterialButton(
          onPressed: () => Navigator.of(context).pushNamed(data.route),
          child: Text(l10n.formEdit),
        ),
      )
    ];
  }
}
