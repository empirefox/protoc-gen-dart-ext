import 'package:flutter/material.dart';

import '../l10n/pgde.l10n.dart';

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
