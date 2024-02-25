import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:firebase_crashlytics/firebase_crashlytics.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:zitch/presentation/views/auth_checker.dart';
import 'package:zitch/theme.dart';

//  This is a FutureProvider that will be used to check whether the firebase has been initialized or not
final firebaseinitializerProvider = FutureProvider<FirebaseApp>((ref) async {
  return await Firebase.initializeApp();
});
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
  runApp(const ProviderScope(child: MyApp()));
}

class MyApp extends ConsumerWidget {
  const MyApp({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final firebaseApp = ref.watch(firebaseinitializerProvider);
    return MaterialApp(
      title: 'Zitch',
      debugShowCheckedModeBanner: false,
      theme: appTheme,
      home: firebaseApp.when(
        data: (app) {
            return const AuthChecker();
          },
        loading: () => const Scaffold(
          body: Center(
            child: CircularProgressIndicator(),
          ),
        ),
        error: (error, stack) => Scaffold(
          body: Center(
            child: Text('Error Occured: $error'),
          ),
        ),
      ),
    );
  }
}
