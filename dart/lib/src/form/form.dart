import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart' show CallOptions, ResponseFuture;
import 'package:protobuf/protobuf.dart';

import '../l10n/pgde.l10n.dart';
import '../pgde/error/error.pb.dart';
import '../validate/validate.dart';

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
  final bool embed;
  final SaveFunc save;
  final ByFieldNumberFunc byNumber;
  final ErrorCoder coder;
  final InputsBuilder builder;

  const FormPage(
      {Key key,
      this.dev = false,
      this.title,
      this.embed,
      this.save,
      this.byNumber,
      this.coder,
      this.builder})
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
      _saving = widget.save(widget.builder.data.editing).then((err) {
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
    _route = null;
  }

  @override
  void dispose() {
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
                      builder: widget.builder,
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
  final InputsBuilder builder;
  final VoidCallback onSave;

  const StatefulForm({Key key, this.builder, this.onSave}) : super(key: key);

  @override
  State<StatefulWidget> createState() => StatefulFormState();
}

class StatefulFormState extends State<StatefulForm> {
  final GlobalKey<FormState> _fbKey = GlobalKey<FormState>();

  GeneratedMessage _draft;

  GeneratedMessage get draft => _draft ??=
      widget.builder.data.editing?.clone() ?? widget.builder.createEmptyDraft();

  bool _needNotSave() => _draft == widget.builder.data.saved;

  void _doSave(BuildContext context) {
    if (!_fbKey.currentState.validate()) return;
    if (_needNotSave()) {
      Navigator.of(context).pop();
      return;
    }
    if (widget.builder.data.isRoot) return setState(widget.onSave);
    widget.builder.data.save(draft);
    Navigator.of(context).pop();
  }

  @override
  Widget build(BuildContext context) {
    final ml10n = MaterialLocalizations.of(context);
    final l10n = PgdeLocalizations.of(context);
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
          child: Column(
            children: widget.builder.buildInputs(context),
          ),
        ),
        Row(
          children: <Widget>[
            Expanded(
              child: MaterialButton(
                color: Theme.of(context).accentColor,
                child: Text(
                  ml10n.okButtonLabel,
                  style: TextStyle(color: Colors.white),
                ),
                onPressed:
                    widget.builder.hasError ? null : () => _doSave(context),
              ),
            ),
            SizedBox(
              width: 20,
            ),
            Expanded(
              child: MaterialButton(
                color: Theme.of(context).accentColor,
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

typedef CreateEmptyDraftFunc<T extends GeneratedMessage> = T Function();

// also act as InputsEmbedMessageOneofBuilder by implementing InputsOneofBuilder
abstract class InputsBuilder<T extends GeneratedMessage>
    implements InputsOneofBuilder {
  static FormFieldValidator catchError(
      VoidCallback fn, List<bool> hasErrorList, int idx) {
    return (v) {
      try {
        fn();
      } catch (e) {
        hasErrorList[idx] = true;
        return e.toString();
      }
      hasErrorList[idx] = false;
      return null;
    };
  }

  List<Widget> buildOneofInputs(BuildContext context, int oo) {
    final oob = oneofBuilders[draft.$_whichOneof(oo)];
    _currentOneofs[oo] = oob;
    return oob.buildInputs(context);
  }

  bool get oneofHasError {
    for (var oo in _currentOneofs) {
      if (oo != null && oo.hasError) return true;
    }
    return false;
  }

  final FormMessageData data;
  final List<InputsOneofBuilder> _currentOneofs;
  InputsBuilder({@required this.data, @required int oneofSize})
      : _currentOneofs = List(oneofSize);

  T _draft;
  T get draft => _draft ??= data.editing?.clone() ?? createEmptyDraft();

  bool _needNotSave() => _draft == data.saved;

  InputDecoration decoration(BuildContext context) =>
      InputDecoration().applyDefaults(Theme.of(context).inputDecorationTheme);

  ValidateInfo<T> vi(BuildContext context) => ValidateInfo<T>(
      MaterialLocalizations.of(context), PgdeLocalizations.of(context), draft);

  CreateEmptyDraftFunc get createEmptyDraft;
  bool get hasError;
  Map<int, InputsOneofBuilder> get oneofBuilders;
  List<Widget> buildInputs(BuildContext context);
}

abstract class InputsOneofBuilder {
  bool get hasError;
  List<Widget> buildInputs(BuildContext context);
}

typedef OneofFieldBuild = Widget Function(
    InputDecoration decoration, FormFieldValidator validator);

class InputsFieldOneofBuilder extends InputsOneofBuilder {
  final InputDecoration decoration;
  final VoidCallback validate;
  final OneofFieldBuild build;

  @override
  bool get hasError => _hasErrorList[0];
  final List<bool> _hasErrorList = [false];

  FormFieldValidator get validator =>
      InputsBuilder.catchError(validate, _hasErrorList, 0);

  InputsFieldOneofBuilder(
      {@required this.decoration,
      @required this.validate,
      @required this.build});

  @override
  List<Widget> buildInputs(BuildContext context) {
    return [build(decoration, validator)];
  }
}

typedef CreateValidatorFunc = Function(BuildContext context, ValidateInfo vi);

typedef CreateDecorationFunc = InputDecoration Function(
    BuildContext context, String errorText);

// non-embed oneof builder
class InputsMessageOneofBuilder extends InputsOneofBuilder {
  // data from childMessage
  final FormMessageData data;
  final CreateEmptyDraftFunc createEmptyDraft;
  final CreateValidatorFunc createValidator;
  final CreateDecorationFunc createDecoration;

  @override
  final bool hasError = false;

  InputsMessageOneofBuilder(
      {@required this.data,
      @required this.createEmptyDraft,
      @required this.createValidator,
      @required this.createDecoration});

  @override
  List<Widget> buildInputs(BuildContext context) {
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
