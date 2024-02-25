import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

ThemeData appTheme = ThemeData(
  textTheme: GoogleFonts.fredokaTextTheme(
    // chnage in fonts as esha messaged
    const TextTheme(
      headlineLarge: TextStyle(),
      headlineMedium: TextStyle(),
      headlineSmall: TextStyle(),
      titleLarge: TextStyle(),
      titleMedium: TextStyle(),
      titleSmall: TextStyle(),
      bodyLarge: TextStyle(),
      bodyMedium: TextStyle(),
      bodySmall: TextStyle(),
    ),
  ),
  useMaterial3: true,
  visualDensity: VisualDensity.adaptivePlatformDensity,
  colorScheme: const ColorScheme(
    background: Colors.black,
    primary: Color.fromRGBO(191, 255, 10, 1),
    brightness: Brightness.dark,
    onPrimary: Colors.black,
    secondary: Color.fromRGBO(66, 72, 246, 1),
    onSecondary: Colors.black,
    error: Colors.red,
    onError: Colors.black,
    onBackground: Colors.white,
    surface: Colors.black,
    onSurface: Colors.white70,
  ),
);
