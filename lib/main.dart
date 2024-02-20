import 'package:flutter/material.dart';
import 'package:zitch/app.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

Future<void> main() async {

  // * Ensures that the Flutter Widgets library is initialized
  WidgetsFlutterBinding.ensureInitialized();

  // * Runs the app
  runApp(
    const ProviderScope(
      child: MyApp(),
    ),
  ); 
}
