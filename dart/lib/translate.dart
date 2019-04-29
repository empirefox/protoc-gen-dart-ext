import 'package:intl/intl.dart';

enum NumName {
  none,
  length,
  mass,
  temperature,
  area,
  volume,
  pressure,
  value,
  bytes,
  characters,
  amount,
  duration,
  date,
  time,
  width,
  height,
}

String numNameString(NumName numName) =>
    numName.toString().substring(numName.runtimeType.toString().length + 1);

class Translator {
  String range_in(String field, String range) {
    return Intl.message(
      '$field must be inside range $range',
      name: 'range_in',
      args: [field, range],
    );
  }

  String range_out(String field, String range) {
    return Intl.message(
      '$field must be outside range $range',
      name: 'range_out',
      args: [field, range],
    );
  }

  String range_in_numName(String field, String range, String numName) {
    return Intl.message(
      '$field ${numName} must be inside range $range',
      name: 'range_in_numName',
      args: [field, range, numName],
      examples: const {
        'numName': <String>["value", "bytes", "characters", "length", "amount"],
      },
    );
  }

  String range_out_numName(String field, String range, String numName) {
    return Intl.message(
      '$field ${numName} must be outside range $range',
      name: 'range_out_numName',
      args: [field, range, numName],
    );
  }

  String in_list(String field, List<dynamic> list) {
    return Intl.message(
      '$field must be one of $list',
      name: 'in_list',
      args: [field, list],
    );
  }

  String not_list(String field, List<dynamic> list) {
    return Intl.message(
      '$field must be not one of $list',
      name: 'not_list',
      args: [field, list],
    );
  }

  String eq_const(String field, String kConst) {
    return Intl.message(
      '$field must equal $kConst',
      name: 'eq_const',
      args: [field, kConst],
    );
  }

  String eq_const_numName(String field, String kConst, String numName) {
    return Intl.message(
      '$field $numName must equal $kConst',
      name: 'eq_const_numName',
      args: [field, kConst, numName],
    );
  }

  String required(String field) {
    return Intl.message(
      '$field is required',
      name: 'required',
      args: [field],
    );
  }

  String numNameOf(NumName numName) {
    switch (numName) {
      case NumName.value:
        return numName_value;
      case NumName.bytes:
        return numName_bytes;
      case NumName.characters:
        return numName_characters;
      case NumName.length:
        return numName_length;
      case NumName.amount:
        return numName_amount;
      default:
        return '';
    }
  }

  String get numName_value {
    return Intl.message(
      'value',
      name: 'numName_value',
    );
  }

  String get numName_bytes {
    return Intl.message(
      'value',
      name: 'numName_bytes',
    );
  }

  String get numName_characters {
    return Intl.message(
      'value',
      name: 'numName_characters',
    );
  }

  String get numName_length {
    return Intl.message(
      'value',
      name: 'numName_length',
    );
  }

  String get numName_amount {
    return Intl.message(
      'value',
      name: 'numName_amount',
    );
  }

  remainingEmailsMessage(int howMany, String userName) => Intl.plural(howMany,
      zero: 'There are no emails left for $userName.',
      one: 'There is $howMany email left for $userName.',
      other: 'There are $howMany emails left for $userName.',
      name: "remainingEmailsMessage",
      args: [howMany, userName],
      desc: "How many emails remain after archiving.",
      examples: const {'howMany': 42, 'userName': 'Fred'});

  notOnlineMessage(String userName, String userGender) =>
      Intl.gender(userGender,
          male: '$userName is unavailable because he is not online.',
          female: '$userName is unavailable because she is not online.',
          other: '$userName is unavailable because they are not online',
          name: "notOnlineMessage",
          args: [userName, userGender],
          desc: "The user is not available to hangout.",
          examples: const {'userGender': 'female', 'userName': 'Alice'});
}
