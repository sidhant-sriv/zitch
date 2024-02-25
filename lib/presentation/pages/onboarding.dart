import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:zitch/presentation/views/auth_checker.dart';

class OnBoardingPage extends StatefulWidget {
  const OnBoardingPage({Key? key}) : super(key: key);

  @override
  State<OnBoardingPage> createState() => _OnBoardingPageState();
}

class _OnBoardingPageState extends State<OnBoardingPage> {
  bool _isLoading = false;
  void loading() {
    setState(() {
      _isLoading = !_isLoading;
    });
  }

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    return Scaffold(
      body: Padding(
        padding: const EdgeInsets.all(15.0),
        child: Consumer(builder: (context, ref, _) {
          final auth = ref.watch(authenticationProvider);
          Future<void> loginWithGoogle() async {
            loading();
            await auth.signInWithGoogle(context).whenComplete(
                  () => auth.authStateChange.listen(
                    (event) async {
                      if (event == null) {
                        loading();
                        return;
                      }
                    },
                  ),
                );
          }

          return Column(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: <Widget>[
              Text('Welcome to zitch!',
                  style: theme.textTheme.titleMedium?.copyWith(
                      color: theme.colorScheme.primary,
                      fontSize: 40,
                      fontWeight: FontWeight.bold)),
              Image.asset('assets/onBoarding.png'),
              Text(
                'Earn points by sharing photos of strays to help NGOs find them and pick them up.',
                style: theme.textTheme.bodyMedium?.copyWith(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                  color: theme.colorScheme.onBackground,
                ),
              ),
              Container(
                padding: const EdgeInsets.only(top: 32.0),
                margin: const EdgeInsets.symmetric(horizontal: 16),
                width: double.infinity,
                child: _isLoading
                    ? const Center(child: CircularProgressIndicator())
                    : MaterialButton(
                        onPressed: loginWithGoogle,
                        textColor: Colors.blue.shade700,
                        textTheme: ButtonTextTheme.primary,
                        minWidth: 100,
                        padding: const EdgeInsets.all(18),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(25),
                          side: BorderSide(color: Colors.blue.shade700),
                        ),
                        child: const Row(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: [
                            Image(
                              image: AssetImage('assets/google_logo.png'),
                              height: 30,
                              width: 30,
                            ),
                            Text(
                              ' Login with Google',
                              style: TextStyle(fontWeight: FontWeight.w600),
                            ),
                          ],
                        ),
                      ),
              ),
            ],
          );
        }),
      ),
    );
  }
}
