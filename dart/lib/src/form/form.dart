import 'package:bot_toast/bot_toast.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart' show CallOptions, ResponseFuture;
import 'package:protobuf/protobuf.dart' show GeneratedMessage;

import '../l10n.dart';
import '../validate.dart';

import 'data.dart';
import 'error.dart';
import 'on_will_pop.dart';
import 'view.dart';

typedef NoNeedAskCallback = bool Function();

typedef SaveFunc<T extends GeneratedMessage> = ResponseFuture<SubmitError>
    Function(T request, {CallOptions options});

class EntryRootCache {
  final Map<String, dynamic> oneof = const {};
  const EntryRootCache();
}

class EntryRoot extends InheritedWidget {
  const EntryRoot({
    Key key,
    @required Widget entry,
  })  : assert(entry != null),
        super(key: key, child: entry);

  final EntryRootCache cache = const EntryRootCache();

  static EntryRoot of(BuildContext context) {
    return context.dependOnInheritedWidgetOfExactType<EntryRoot>();
  }

  @override
  bool updateShouldNotify(EntryRoot old) => false;
}

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

  Future<SubmitError> _saving;

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
          builder: (BuildContext context, AsyncSnapshot<SubmitError> snapshot) {
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
                final children = <Widget>[
                  StatefulForm(
                    inputs: widget.inputs,
                    onSave: () => _doSave(context),
                  )
                ];
                if (err != null) {
                  if (err.path == null || err.path.length == 0)
                    children.add(OperateErrorText(
                      err.code,
                      widget.coder,
                      err.message,
                    ));
                }
                return Column(children: children);
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

  GeneratedMessage get draft =>
      _draft ??= widget.inputs.data.draft(widget.inputs.createEmptyDraft);

  bool _needNotSave() => _draft == widget.inputs.data.saved;

  void _doSave(BuildContext context, PgdeLocalizations l10n) {
    if (!_fbKey.currentState.validate()) {
      BotToast.showText(text: l10n.formValidateFailed);
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

abstract class InputsBase<T extends GeneratedMessage> {
  String get path;
  ErrorCoderGetter get coder;
  PgdeLocalizations get pgdeL10n;
  List<Widget> build(BuildContext context);
}

// wrap leaf inputs to one widget
class InputsWidget<T extends GeneratedMessage> extends StatelessWidget {
  final InputsBase<T> inputs;
  final SubmitError err;
  const InputsWidget(this.inputs, this.err, {Key key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    final color = Theme.of(context).errorColor ?? Colors.red;
    final List<Widget> list = inputs.build(context);
    if (err?.path == inputs.path) {
      list.insert(
          0,
          Text(
            err.text(inputs.pgdeL10n, inputs.coder.of(T)),
            style: TextStyle(color: color),
          ));
      return Container(
        child: Column(children: list),
        decoration: BoxDecoration(border: Border.all(color: color)),
      );
    }
    return Column(children: list);
  }
}

// for every leaf
abstract class Inputs<T extends GeneratedMessage> implements InputsBase<T> {
  static String catchError(VoidCallback fn) {
    try {
      fn();
    } catch (e) {
      return e.toString();
    }
    return null;
  }

  final String path;
  final FormMessageData data;
  final SubmitError err;
  final ErrorCoderGetter coder;
  final PgdeLocalizations pgdeL10n;
  Inputs(this.path, this.data, this.err, this.coder, this.pgdeL10n);

  T _draft;
  T get draft => _draft ??= data.draft(createEmptyDraft);

  bool _needNotSave() => _draft == data.saved;

  InputDecoration decoration(BuildContext context) =>
      InputDecoration().applyDefaults(Theme.of(context).inputDecorationTheme);

  ValidateInfo<T> vi(BuildContext context) => ValidateInfo<T>(
      MaterialLocalizations.of(context), PgdeLocalizations.of(context), draft);

  CreateEmptyDraftFunc<T> get createEmptyDraft;

  FormFieldValidator validator(int tag, VoidCallback fn) {
    if (err.path == '${this.path}/$tag') {
      if (err.when == null) err.when = data.editing.getField(tag);
      return (v) {
        if (err.when == v)
          return errorText(pgdeL10n, coder.of(T), err.code, err.message);
        return Inputs.catchError(fn);
      };
    }
    return (v) => Inputs.catchError(fn);
  }

  @override
  List<Widget> build(BuildContext context);
}

class SelectSource extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return ListView(
        // TODO
        );
  }
}

class InfiniteSelectSource extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return InfiniteListView(
      nextList: (page) {
        return size;
      },
    );
  }
}

class SelectDest extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    throw UnimplementedError();
  }
}

// for select one/many field
class ChildFieldEntryWidget extends StatelessWidget {
  // data from childMessage
  final FormMessageData data;
  final InputDecoration decoration;

  ChildFieldEntryWidget({@required this.data, @required this.decoration});

  @override
  Widget build(BuildContext context) {
    final l10n = PgdeLocalizations.of(context);
    return InputDecorator(
      decoration:
          decoration.applyDefaults(Theme.of(context).inputDecorationTheme),
      child: MaterialButton(
        onPressed: () => Navigator.of(context).pushNamed(data.route),
        child: Text(l10n.formEdit),
      ),
    );
  }
}
