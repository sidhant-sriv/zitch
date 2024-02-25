import 'dart:io';

import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';

class CameraView extends StatefulWidget {
  const CameraView({super.key});

  @override
  State<CameraView> createState() => _CameraViewState();
}

class _CameraViewState extends State<CameraView> {
  File? takenImage;

  void takePicture(ImageSource source) async {
    final ImagePicker picker = ImagePicker();
    final XFile? image = await picker.pickImage(source: source, maxWidth: 600);
    if (image == null) {
      return;
    }
    setState(() {
      takenImage = File(image.path);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Container(
          alignment: Alignment.center,
          decoration: BoxDecoration(
              border: Border.all(
            color: Theme.of(context).colorScheme.secondary.withOpacity(0.5),
            width: 3.0,
          )),
          height: MediaQuery.of(context).size.height * 0.6,
          width: double.infinity,
          child: takenImage != null
              ? Image.file(
                  takenImage!,
                  fit: BoxFit.contain,
                  width: double.infinity,
                )
              : Text(
                  'No Image Taken',
                  style: TextStyle(
                    color:
                        Theme.of(context).colorScheme.primary.withOpacity(0.5),
                  ),
                ),
        ),
        SizedBox(
          height: MediaQuery.of(context).size.height * 0.05,
        ),
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          children: [
            FloatingActionButton(
              onPressed: () => takePicture(ImageSource.camera),
              child: const Icon(Icons.camera_alt),
            ),
            FloatingActionButton(
              onPressed: () => takePicture(ImageSource.gallery),
              child: const Icon(Icons.file_open),
            ),
          ],
        ),
      ],
    );
  }
}
