import 'package:flutter/material.dart';

import '../l10n.dart';

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
