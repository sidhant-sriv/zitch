import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'presentation/views/home_view.dart';

final GlobalKey<NavigatorState> navigatorKey = GlobalKey<NavigatorState>();

final router = GoRouter(
  navigatorKey: navigatorKey,
  debugLogDiagnostics: true,
  initialLocation: '/',
  routes: [
    GoRoute(
      path: '/',
      builder: (context, state) => const HomeView(),
    ),
  ],
);
