import 'package:flutter/material.dart';
import 'package:monggo/entity/user.dart';
import 'package:monggo/presentation/auth/components/header.dart';
import 'package:monggo/presentation/home/components/messages_list.dart';

class Home extends StatelessWidget {
  List<ChatUsers> chatUsers = [
    ChatUsers(
        name: "085747672701",
        messageText: "Awesome Setup",
        imageURL: "images/userImage1.jpeg",
        time: "18:50"),
    ChatUsers(
        name: "085747672711",
        messageText: "Awesome Setup",
        imageURL: "images/userImage1.jpeg",
        time: "19:50"),
    ChatUsers(
        name: "085747672791",
        messageText: "Awesome Setup",
        imageURL: "images/userImage1.jpeg",
        time: "18:20"),
    ChatUsers(
        name: "08574767211",
        messageText: "Awesome Setup",
        imageURL: "images/userImage1.jpeg",
        time: "18:10"),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Color(0xFF497fff),
      body: SafeArea(
          child: Stack(
        children: [
          Container(
            height: MediaQuery.of(context).size.height,
            width: MediaQuery.of(context).size.width,
            color: Color(0xFF497fff),
          ),
          CustomHeader(
            text: 'MONGGO SERVER.',
            onTap: () {},
          ),
          Positioned(
            top: MediaQuery.of(context).size.height * 0.10,
            child: Container(
              height: MediaQuery.of(context).size.height * 0.9,
              width: MediaQuery.of(context).size.width,
              decoration: const BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.only(
                      topLeft: Radius.circular(25),
                      topRight: Radius.circular(25))),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  SizedBox(
                    height: 25,
                  ),
                  Center(
                    child: Container(
                      width: MediaQuery.of(context).size.width * 0.8,
                      height: 50,
                      decoration: BoxDecoration(
                          color: Color(0xFFF2F2F2),
                          borderRadius: BorderRadius.all(Radius.circular(20))),
                      child: Row(
                        children: [
                          Container(
                            margin: const EdgeInsets.only(left: 5, right: 2),
                            width: 100,
                            height: 40,
                            decoration: BoxDecoration(
                                borderRadius:
                                    BorderRadius.all(Radius.circular(15))),
                          ),
                          Container(
                            margin: const EdgeInsets.only(left: 5, right: 2),
                            width: 100,
                            height: 40,
                            decoration: BoxDecoration(
                                color: Colors.white,
                                borderRadius:
                                    BorderRadius.all(Radius.circular(15))),
                          ),
                          Container(
                            margin: const EdgeInsets.only(left: 5, right: 2),
                            width: 100,
                            height: 40,
                            decoration: BoxDecoration(
                                color: Color(0xFFF2F2F2),
                                borderRadius:
                                    BorderRadius.all(Radius.circular(15))),
                          ),
                        ],
                      ),
                    ),
                  ),
                  ListView.builder(
                    itemCount: chatUsers.length,
                    shrinkWrap: true,
                    padding: EdgeInsets.only(top: 16),
                    physics: NeverScrollableScrollPhysics(),
                    itemBuilder: (context, index) {
                      return Conversation(
                        name: chatUsers[index].name,
                        messagesText: chatUsers[index].messageText,
                        imageUrl: chatUsers[index].imageURL,
                        time: chatUsers[index].time,
                        isMessageRead:
                            (index == 0 || index == 3) ? true : false,
                      );
                    },
                  ),
                ],
              ),
            ),
          ),
        ],
      )),
    );
  }
}
