import 'package:flutter/material.dart';

import '../models/community.dart';

class Card extends StatefulWidget {
  const Card({Key? key, required this.cur}) : super(key: key);
  final Community cur;
  @override
  State<Card> createState() => _CardState();
}

class _CardState extends State<Card> {
  @override
  Widget build(BuildContext context) {
    return Padding(
        padding: const EdgeInsets.only(left: 5, right: 5, top: 10),
        child: Container(
          child: Padding(
            child: Text(widget.cur.name!),
            padding: const EdgeInsets.only(left: 20),
          ),
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(15),
            color: Colors.lightGreen,
          ),
        ));
  }
}
