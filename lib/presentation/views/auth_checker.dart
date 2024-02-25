import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:zitch/presentation/pages/home_page.dart';
import 'package:zitch/presentation/pages/onboarding.dart';
import 'package:zitch/repository/auth_repository.dart';

final authenticationProvider = Provider<Authentication>((ref) {
  return Authentication();
});

final authStateProvider = StreamProvider<User?>((ref) {
  return ref.read(authenticationProvider).authStateChange;
});

class AuthChecker extends ConsumerWidget {
  const AuthChecker({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final authState = ref.watch(authStateProvider);
    return authState.when(
        data: (data) {
          print("User data $data");
          if (data != null) return const HomePage();
          return const OnBoardingPage();
        },
        loading: () => const CircularProgressIndicator(),
        error: (e, trace) => Text('Error Occured: $e'));
  }
}
