import 'package:flutter/material.dart';

import '../l10n.dart';

const justError = 1;
const withCode = 2;
const withCoder = 3;
const withDetail = 4;
int errorMode = withCoder;

typedef ByFieldNumberFunc = String Function(int fieldNumber);

abstract class ErrorCoderGetter {
  ErrorCoder of(Type type);
}

abstract class ErrorCoder {
  String $byErrorCode(int number);
}

String errorText(
    PgdeLocalizations l10n, ErrorCoder coder, int code, String message) {
  switch (errorMode) {
    case justError:
      return l10n.formOperateFailed;
    case withCode:
      return l10n.formOperateFailedWithCode(code);
    case withDetail:
      if (message.isEmpty) return coder.$byErrorCode(code);
      return coder.$byErrorCode(code) + ': ' + message;
    default:
      return coder.$byErrorCode(code);
  }
}

class SubmitError {
  final String path;
  final int code;
  final String message;

  String when;

  SubmitError(this.path, this.code, this.message);

  String text(PgdeLocalizations l10n, ErrorCoder coder) =>
      errorText(l10n, coder, code, message);
}

class OperateErrorText extends StatelessWidget {
  final int code;
  final ErrorCoder coder;
  final String message;

  const OperateErrorText(this.code, this.coder, this.message, {Key key})
      : super(key: key);

  @override
  Widget build(BuildContext context) =>
      Text(errorText(PgdeLocalizations.of(context), coder, code, message));
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
