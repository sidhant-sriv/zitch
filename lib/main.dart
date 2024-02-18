import 'package:flutter/material.dart';
import 'package:zitch/app.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:firebase_crashlytics/firebase_crashlytics.dart';
import 'package:flutter/foundation.dart';

Future<void> main() async {
  // * Ensures that the Flutter Widgets library is initialized
  WidgetsFlutterBinding.ensureInitialized();

  // * Initializes Firebase
  await Firebase.initializeApp(
      options: const FirebaseOptions(
    appId: '1:586074781896:android:c1f321d1c43246ccd78dfd',
    messagingSenderId: '586074781896',
    projectId: 'zitch-5f2b1',
    apiKey: 'AIzaSyATx3dOxGOanlBrf50qcxKNnx6bkepnBME',
  ));

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
