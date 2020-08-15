import 'package:flutter/material.dart';
import 'package:pgde/pgde.dart' as $pgde;
import 'package:protobuf/protobuf.dart';

import '../pgde/error/error.pb.dart' as $error;
import 'config.pb.dart';
import 'config.l10n.dart';
import 'config.validate.dart';
import 'config.zero.dart';

class AdpRouterInputs extends $pgde.Inputs<AdpRouter> {
  AdpRouterInputs(
    String path,
    $pgde.FormMessageData data,
    $pgde.SubmitError err,
    $pgde.ErrorCoderGetter coder,
    $pgde.PgdeLocalizations pgdeL10n,
  ) : super(path, data, err, coder, pgdeL10n);

  @override
  List<Widget> build(BuildContext context) {
    // TODO: implement build
    return null;
  }

  @override
  // TODO: implement createEmptyDraft
  get createEmptyDraft => null;
}

// message example
class LogInputs extends $pgde.Inputs<Log> {
  @override
  final $pgde.CreateEmptyDraftFunc<Log> createEmptyDraft;

  LogInputs(String path, $pgde.FormMessageData data, $pgde.SubmitError err,
      $pgde.ErrorCoderGetter coder, $pgde.PgdeLocalizations pgdeL10n)
      : createEmptyDraft = zeroLog.create,
        super(path, data, err, coder, pgdeL10n);

  @override
  List<Widget> build(BuildContext context) {
    final l10n = ConfigLocalizations.of(context);
    final v = LogValidator(context, vi(context));
    final root = $pgde.EntryRoot.of(context);

    final draft2 = RouterUseView_Id();

    final f3 = path + '/3';

    return <Widget>[
      // normal_field_not_in_oneof
      TextFormField(
        decoration: InputDecoration(
          labelText: l10n.LogTargetLabel,
          hintText: l10n.LogTargetHint,
        ),
        initialValue: draft.target,
        onChanged: (n) => draft.target = n,
        autovalidate: true,
        validator: validator(1, v.assertField_target),
      ),
      // oneof, TODO: cache oneof from root entry?
      LogRouterTabs([
        Column(
          children: [
            const Icon(Icons.insert_comment),
            Text(l10n.IPNetRouter),
          ],
        ),
      ], [
        // normal_field_in_oneof, just like normal field
        TextFormField(
          decoration: InputDecoration(
            labelText: l10n.LogTargetLabel,
            hintText: l10n.LogTargetHint,
          ),
          initialValue: draft2.whichIs() == RouterUseView_Id_Is.adp
              ? draft.target
              : root.cache.oneof[f3],
          onChanged: (n) {
            draft.target = n;
            root.cache.oneof[f3] = n;
          },
          autovalidate: true,
          validator: validator(1, v.assertField_target),
        ),
      ]),
      // embed
      $pgde.InputsWidget<AdpRouter>(
        AdpRouterInputs(
          path + '/5',
          data.childMessage(5),
          err,
          coder,
          pgdeL10n,
        ),
        err,
      ),
      // TODO: move to entry field
      $pgde.ChildFieldEntryWidget(
        data: data.childMessage(4),
        decoration: InputDecoration(
          labelText: l10n.IPNetRouterLabel,
          hintText: l10n.IPNetRouterHint,
          errorText: $pgde.Inputs.catchError(validator.assertField_ipnet),
        ),
      ),
      // TODO select_one
      // TODO select_many
    ];
  }
}

// list example
class LogInputsIpfsServersInput extends StatelessWidget {
  final FormMessageData data;
  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return null;
  }
}

// oneof example
class LogRouterTabs extends StatelessWidget {
  final List<Widget> tabBarList;
  final List<Widget> tabBarViewList;

  LogRouterTabs(this.tabBarList, this.tabBarViewList, {Key key})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
        length: tabBarList.length,
        child: Column(
          children: <Widget>[
            Container(
              width: double.infinity,
              height: 80,
              padding: EdgeInsets.fromLTRB(20, 24, 0, 0),
              alignment: Alignment.centerLeft,
              color: Colors.black,
              child: TabBar(
                isScrollable: true,
                indicatorColor: Colors.red,
                indicatorSize: TabBarIndicatorSize.label,
                unselectedLabelColor: Colors.white,
                unselectedLabelStyle: TextStyle(fontSize: 18),
                labelColor: Colors.red,
                labelStyle: TextStyle(fontSize: 20),
                tabs: tabBarList,
              ),
            ),
            Expanded(child: TabBarView(children: tabBarViewList)),
          ],
        ));
  }
}
