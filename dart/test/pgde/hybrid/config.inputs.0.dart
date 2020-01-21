import 'package:flutter/material.dart';
import 'package:flutter_form_builder/flutter_form_builder.dart';
import 'package:pgde/pgde.dart' as $pgde;
import 'package:protobuf/protobuf.dart';

import '../google/protobuf/empty.pb.dart';
import '../pgde/error/error.pb.dart';
import 'config.pb.dart';
import 'config.l10n.dart';
import 'config.rpc.pbgrpc.dart';
import 'config.validate.dart';

class AdpRouter extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return null;
  }
}

class IPNetRouter extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return null;
  }
}

class RouterItemPage extends StatefulWidget {
  final bool dev;
  final bool embed;
  final RouterItemServiceClient client;
  final $pgde.ErrorCoder coder;
  final $pgde.FormMessageData<String> data;

  const RouterItemPage(
      {Key key,
      this.dev = false,
      this.embed,
      this.client,
      this.coder,
      this.data})
      : super(key: key);

  @override
  RouterItemState createState() {
    return RouterItemState();
  }
}

class RouterItemState extends State<RouterItemPage> {
  final GlobalKey<FormBuilderState> _fbKey = GlobalKey<FormBuilderState>();

  String _route;
  String route(BuildContext context) =>
      _route ??= ModalRoute.of(context).settings.name;

  $pgde.FormMessageData<String> get data => widget.data;

  Future<BackendError> _saving;

  void _doSave(BuildContext context) {
    setState(() {
      _saving = widget.client.setRouterItem(widget.data.editing).then((err) {
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
    final l10n = ConfigLocalizations.of(context);
    return Scaffold(
      appBar: AppBar(
        title: Text(l10n.RouterItem + (widget.dev ? '-${route(context)}' : '')),
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
                  return $pgde.GrpcErrorRetry(
                    err: snapshot.error,
                    onRetry: () => _doSave(context),
                  );

                final err = snapshot.data;
                return Column(
                  children: <Widget>[
                    RouterItemForm(
                      data: widget.data,
                      onSave: () => _doSave(context),
                    ),
                    err == null
                        ? Container()
                        : $pgde.ErrorText(
                            err.fieldNumber,
                            l10n.$ofRouterItem,
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

  List<Widget> buildInputs(BuildContext context, RouterItem draft) {
    return <Widget>[];
  }
}

class RouterItemForm extends StatefulWidget {
  final $pgde.FormMessageData<String> data;
  final VoidCallback onSave;

  const RouterItemForm({Key key, this.data, this.onSave}) : super(key: key);

  @override
  State<StatefulWidget> createState() => RouterItemFormSate();
}

class RouterItemFormSate extends State<RouterItemForm> {
  final GlobalKey<FormBuilderState> _fbKey = GlobalKey<FormBuilderState>();

  GeneratedMessage _draft;
  GeneratedMessage get draft => _draft ??= (widget.data.editing == null
      ? createEmptyDraft()
      : widget.data.editing.clone());

  bool _needNotSave() => _draft == widget.data.saved;

  void _doSave(BuildContext context) {
    if (!_fbKey.currentState.validate()) return;
    if (_needNotSave()) {
      Navigator.of(context).pop();
      return;
    }
    if (widget.data.isRoot) return setState(widget.onSave);
    widget.data.save(draft);
    Navigator.of(context).pop();
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: <Widget>[
        Form(
          key: _fbKey,
          // autovalidate: true,
          // readonly: true,
          onWillPop: () => $pgde.onStayOrWillPop(
            context,
            () => Navigator.of(context).pop(),
            noNeedAsk: _needNotSave(),
          ),
          child: Column(
            children: buildInputs(context, draft, _hasErrorList),
          ),
        ),
        Row(
          children: <Widget>[
            Expanded(
              child: MaterialButton(
                color: Theme.of(context).accentColor,
                child: Text(
                  "Submit",
                  style: TextStyle(color: Colors.white),
                ),
                onPressed: hasError ? null : () => _doSave(context),
              ),
            ),
            SizedBox(
              width: 20,
            ),
            Expanded(
              child: MaterialButton(
                color: Theme.of(context).accentColor,
                child: Text(
                  "Reset",
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

  bool get hasError => _hasErrorList.any((n) => n);
  final List<bool> _hasErrorList = List<bool>.filled(1, false);

  FormFieldValidator _catch(VoidCallback fn, List<bool> hasErrorList, int idx) {
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

  GeneratedMessage createEmptyDraft() => RouterItem();

  List<Widget> buildInputs(
      BuildContext context, GeneratedMessage draft, List<bool> hasErrorList) {
    final data = draft as RouterItem;
    final vi = $pgde.ValidateInfo(MaterialLocalizations.of(context),
        $pgde.PgdeLocalizations.of(context), null, data);
    final validator = RouterItemValidator(context, vi);
    return <Widget>[
      TextFormField(
        initialValue: data.name,
        onChanged: (n) => data.name = n,
        autovalidate: true,
        validator: _catch(validator.assertField_name, hasErrorList, 0),
      )
    ];
  }
}
