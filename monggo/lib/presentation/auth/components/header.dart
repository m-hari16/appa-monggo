import 'package:flutter/material.dart';

class CustomHeader extends StatelessWidget {
  final String text;
  final Function()? onTap;
  const CustomHeader({Key? key, required this.text, required this.onTap})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(top: 16, left: 16),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          const SizedBox(
            width: 16,
          ),
          Text(
            text,
            style: TextStyle(
                color: Color(0xFFF8F9FA),
                fontSize: 28,
                fontWeight: FontWeight.w700),
          )
        ],
      ),
    );
  }
}
