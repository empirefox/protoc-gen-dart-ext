import 'dart:io';

import 'package:args/args.dart';

import '../lib/protos/exports/exports.pb.dart';

const input = 'input';

void main(List<String> args) {
  final parser = ArgParser()
    ..addOption(input, abbr: 'i', help: 'the required input file name');
  final argResults = parser.parse(args);
  final outfile = argResults[input];
  if (outfile == null) {
    print(parser.usage);
    exit(-1);
  }

  final contents = File(outfile).readAsBytesSync();
  final e = Exports.fromBuffer(contents);
  print(e);
}
