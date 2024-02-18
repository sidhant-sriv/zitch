import 'package:flutter/material.dart';
import 'package:zitch/app.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:firebase_crashlytics/firebase_crashlytics.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

Future<void> main() async {
  // * Ensures that the Flutter Widgets library is initialized
  WidgetsFlutterBinding.ensureInitialized();

  // * Loads the .env file
  await dotenv.load(fileName: '.env');

  // * Initializes Firebase
  await Firebase.initializeApp(
    options: FirebaseOptions(
      appId: dotenv.env['APP_ID'] ?? '',
      messagingSenderId: dotenv.env['MESSAGING_ID'] ?? '',
      projectId: dotenv.env['PROJECT_ID'] ?? '',
      apiKey: dotenv.env['API_KEY'] ?? '',
    ),
  );

  // * Crashlytics Initialization
  FlutterError.onError = (errorDetails) {
    FirebaseCrashlytics.instance.recordFlutterFatalError(errorDetails);
  };
  PlatformDispatcher.instance.onError = (error, stack) {
    FirebaseCrashlytics.instance.recordError(error, stack, fatal: true);
    return true;
  };

  // * Runs the app
  runApp(
    const ProviderScope(
      child: MyApp(),
    ),
  );
}
