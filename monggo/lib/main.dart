import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:monggo/app/routes.dart';

void main() {
  runApp(Myapp());
}

class Myapp extends StatefulWidget {
  const Myapp({Key? key}) : super(key: key);

  @override
  State<Myapp> createState() => _MyappState();
}

class _MyappState extends State<Myapp> {
  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      title: "Monggo app",
      getPages: Routes.register,
      initialRoute: "/login",
    );
  }
}
