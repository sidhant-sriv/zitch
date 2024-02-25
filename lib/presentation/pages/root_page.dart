import 'package:flutter/material.dart';
import 'package:zitch/presentation/views/camera_view.dart';
import 'package:zitch/presentation/views/community_view.dart';
import 'package:zitch/presentation/views/maps_view.dart';
import 'package:zitch/presentation/views/profile_view.dart';

class RootPage extends StatefulWidget {
  const RootPage({
    super.key,
  });

  @override
  State<RootPage> createState() => _RootPageState();
}

class _RootPageState extends State<RootPage> {
  int _currentIndex = 0;

  @override
  Widget build(BuildContext context) {
    final List<Widget> _children = [
      const CameraView(),
      const CommunityView(),
      const MapsView(),
      const ProfileView(),
    ];
    return Scaffold(
      appBar: AppBar(
        toolbarHeight: 100,
        elevation: 0,
        surfaceTintColor: Colors.transparent,
        title: Center(
          child: Padding(
              padding: const EdgeInsets.all(12.0),
              child: Image.asset(
                'assets/zitch_icon.png',
                height: 40,
              )),
        ),
      ),
      body: _children[_currentIndex],
      bottomNavigationBar: BottomNavigationBar(
        useLegacyColorScheme: false,
        currentIndex: _currentIndex,
        onTap: (index) {
          setState(() {
            _currentIndex = index;
          });
        },
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.camera_alt_outlined),
            activeIcon: Icon(Icons.camera_alt_rounded),
            label: 'Camera',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.people_alt_outlined),
            activeIcon: Icon(Icons.people_outline_rounded),
            label: 'Community',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.map_outlined),
            activeIcon: Icon(Icons.map_rounded),
            label: 'Info',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.pets_outlined),
            activeIcon: Icon(Icons.pets_rounded),
            label: 'Profile',
          ),
        ],
      ),
    );
  }
}
