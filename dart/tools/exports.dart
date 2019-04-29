import 'dart:io';

import 'package:args/args.dart';
import '../lib/exports/register.dart';

import '../lib/pgde.exports.dart' as $pkga;

void register() {
  $pkga.register();
}

const output = 'output';
void main(List<String> args) {
  final parser = ArgParser()
    ..addOption(output, abbr: 'o', help: 'the required output file name');
  final argResults = parser.parse(args);
  final outfile = argResults[output];
  if (outfile == null) {
    print(parser.usage);
    exit(-1);
  }

  register();
  writeToFile(outfile);
}
