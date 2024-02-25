import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:zitch/presentation/pages/onboarding.dart';
import 'presentation/views/community_view.dart';
import 'presentation/views/camera_view.dart';
import 'presentation/views/maps_view.dart';
import 'presentation/views/profile_view.dart';

final GlobalKey<NavigatorState> navigatorKey = GlobalKey<NavigatorState>();

final router = GoRouter(
  navigatorKey: navigatorKey,
  debugLogDiagnostics: true,
  initialLocation: '/board',
  routes: <RouteBase>[
    GoRoute(
      path: '/board',
      builder: (context, state) => const OnBoardingPage(),
    ),
    ShellRoute(
      routes: [
        GoRoute(
          path: '/home',
          builder: (context, state) => Container(
            decoration: BoxDecoration(
              color: Theme.of(context).colorScheme.background,
            ),
            width: MediaQuery.of(context).size.width,
            height: MediaQuery.of(context).size.height,
            child: const CameraView(),
          ),
        ),
        GoRoute(
          path: '/community',
          builder: (context, state) => Container(
            decoration: BoxDecoration(
              color: Theme.of(context).colorScheme.background,
            ),
            width: MediaQuery.of(context).size.width,
            height: MediaQuery.of(context).size.height,
            child: const CommunityView(),
          ),
        ),
        GoRoute(
          path: '/maps',
          builder: (context, state) => Container(
            decoration: BoxDecoration(
              color: Theme.of(context).colorScheme.background,
            ),
            width: MediaQuery.of(context).size.width,
            height: MediaQuery.of(context).size.height,
            child: const MapsView(),
          ),
        ),
        GoRoute(
          path: '/profile',
          builder: (context, state) => Container(
            decoration: BoxDecoration(
              color: Theme.of(context).colorScheme.background,
            ),
            width: MediaQuery.of(context).size.width,
            height: MediaQuery.of(context).size.height,
            child: const ProfileView(),
          ),
        ),     
      ],
    ),
  ],
);
