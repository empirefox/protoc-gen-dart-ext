import 'package:intl/intl.dart';

// Only used to generate arb!
// make gen_validate_arb
// make rewrite_validate_arb
class Translator {
  const Translator._();

  String get noMoreText => Intl.message(
        'No More',
        desc: 'show user there is no more item',
        name: 'noMoreText',
      );
  String get retry => Intl.message(
        'Retry',
        desc: 'retry when errors occur',
        name: 'retry',
      );
  String get backAlertTitle => Intl.message(
        'Alert',
        desc: 'title of alert which warns changes exist',
        name: 'backAlertTitle',
      );
  String get backAlertContent => Intl.message(
        'Changes exist!',
        desc: 'content of alert which warns changes exist',
        name: 'backAlertContent',
      );
  String get backAlertStay => Intl.message(
        'Stay',
        desc: 'choose to stay when there are changes',
        name: 'backAlertStay',
      );
  String get backAlertGoBack => Intl.message(
        'Go back',
        desc: 'choose to go back when there are changes',
        name: 'backAlertGoBack',
      );
  String get reset => Intl.message(
        'Reset',
        desc: 'reset to default',
        name: 'reset',
      );
  String get edit => Intl.message(
        'Edit',
        desc: 'edit oneof message',
        name: 'edit',
      );
  String get validateFailed => Intl.message(
        'Form validation failed!',
        desc: 'form validate failed',
        name: 'validateFailed',
      );
  String get operateFailed => Intl.message(
        'Operate Failed',
        desc: 'operate failed',
        name: 'operateFailed',
      );
  String operateFailedWithCode(int code) => Intl.message(
        'Operate Failed with code [$code]',
        desc: 'operate failed with code',
        name: 'operateFailedWithCode',
        args: [code],
      );
}
