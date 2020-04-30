import 'package:flutter/material.dart';
import 'package:pgde/pgde.dart' as $pgde;
import 'package:protobuf/protobuf.dart';

import 'config.pb.dart';
import 'config.l10n.dart';
import 'config.validate.dart';
import 'config.zero.dart';

class AdpRouterInputs extends $pgde.Inputs<AdpRouter> {
  AdpRouterInputs($pgde.FormMessageData data) : super(data);

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
class RouterItemInputs extends $pgde.Inputs<RouterItem> {
  @override
  final $pgde.CreateEmptyDraftFunc<RouterItem> createEmptyDraft;

  RouterItemInputs($pgde.FormMessageData data)
      : createEmptyDraft = zeroRouterItem.create,
        super(data);

  @override
  List<Widget> build(BuildContext context) {
    final l10n = ConfigLocalizations.of(context);
    final validator = RouterItemValidator(context, vi(context));

    return <Widget>[
      // field_not_in_oneof
      TextFormField(
        decoration: InputDecoration(
          labelText: l10n.RouterItemNameLabel,
          hintText: l10n.RouterItemNameHint,
        ),
        initialValue: draft.name,
        onChanged: (n) => draft.name = n,
        autovalidate: true,
        validator: $pgde.Inputs.catcher(validator.assertField_name),
      ),
      // oneof
      RouterItemRouterTabs([
        const Icon(Icons.insert_comment),
        Text(l10n.AdpRouter),
        Column(
          children: [
            const Icon(Icons.insert_comment),
            Text(l10n.IPNetRouter),
          ],
        ),
      ], [
        // field_in_oneof
        TextFormField(
          decoration: InputDecoration(
            labelText: l10n.RouterItemNameLabel,
            hintText: l10n.RouterItemNameHint,
          ),
          initialValue: draft.name,
          onChanged: (n) => draft.name = n,
          autovalidate: true,
          validator: $pgde.Inputs.catcher(validator.assertField_name),
        ),
        // embed
        $pgde.InputsWidget<AdpRouter>(AdpRouterInputs(data.childMessage(3))),
        // child_form_entry
        $pgde.ChildFieldEntryWidget(
          data: data.childMessage(4),
          decoration: InputDecoration(
            labelText: l10n.IPNetRouterLabel,
            hintText: l10n.IPNetRouterHint,
            errorText: $pgde.Inputs.catchError(validator.assertField_ipnet),
          ),
        ),
      ]),
      // TODO foreign_link
      // TODO list_entry
      // TODO map_entry
    ];
  }
}

// list example
class RouterItemInputsIpfsServersInput extends StatelessWidget {
  final FormMessageData data;
  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return null;
  }
}

// oneof example
class RouterItemRouterTabs extends StatelessWidget {
  final List<Widget> tabBarList;
  final List<Widget> tabBarViewList;

  RouterItemRouterTabs(this.tabBarList, this.tabBarViewList, {Key key})
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
