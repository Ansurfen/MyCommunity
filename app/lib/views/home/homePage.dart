import 'dart:convert';

import 'package:dio/dio.dart' as dio;
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:my_community/api/community.dart';
import 'package:my_community/components/card.dart' as mine;
import 'package:my_community/components/drawer.dart';
import 'package:my_community/models/community.dart';
import 'package:my_community/store/conf.dart';
import 'package:my_community/store/user.dart';
import 'package:my_community/views/home/homeController.dart';

class HomePage extends StatefulWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  int _selectedIndex = 0;
  HomeController homeController = Get.find<HomeController>();

  @override
  void initState() {
    super.initState();
    homeController.pages
      ..add(const Search())
      ..add(const Backed())
      ..add(const About());
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("社团云平台"),
        actions: <Widget>[
          IconButton(onPressed: () {}, icon: const Icon(Icons.share))
        ],
      ),
      drawer: getConf().login ? LoggedInDrawer() : const UnLoginDrawer(),
      bottomNavigationBar: BottomNavigationBar(
        items: const <BottomNavigationBarItem>[
          BottomNavigationBarItem(icon: Icon(Icons.home), label: '社团'),
          BottomNavigationBarItem(icon: Icon(Icons.business), label: '后台'),
          BottomNavigationBarItem(icon: Icon(Icons.school), label: '关于'),
        ],
        currentIndex: _selectedIndex,
        fixedColor: Colors.blue,
        onTap: (index) {
          setState(() {
            _selectedIndex = index;
          });
        },
      ),
      body: homeController.pages[_selectedIndex],
      floatingActionButton:
          FloatingActionButton(child: const Icon(Icons.add), onPressed: _onAdd),
    );
  }

  void _onAdd() async {

  }
}

class Search extends StatefulWidget {
  const Search({Key? key}) : super(key: key);

  @override
  State<Search> createState() => _SearchState();
}

class _SearchState extends State<Search> {
  final GlobalKey _key = GlobalKey<State>();
  final communities = Rx<List<Map<String, dynamic>>>([]);
  String? input;

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(top: 25, left: 25, right: 25),
      child: Column(
        children: [
          Row(
            children: [
              Flexible(
                child: TextFormField(
                  key: _key,
                  onTap: () async {
                    dio.FormData _formData =
                        dio.FormData.fromMap({"cname": input});
                    dio.Response res = await searchCommunity(_formData);
                    communities.value =
                        (json.decode(res.data['data']['data']) as List)
                            .map((e) => e as Map<String, dynamic>)
                            .toList();
                    print(res.data['data']['data']);
                  },
                  onChanged: (v) => input = v,
                  cursorColor: Colors.grey,
                  decoration: InputDecoration(
                      fillColor: Colors.white,
                      filled: true,
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(10),
                        borderSide: BorderSide.none,
                      ),
                      hintText: 'Search',
                      hintStyle:
                          const TextStyle(color: Colors.grey, fontSize: 18),
                      prefixIcon: Container(
                        padding: const EdgeInsets.all(15),
                        child: const Icon(Icons.search),
                        width: 18,
                      )),
                ),
                flex: 1,
              )
            ],
          ),
          Expanded(
              child: Center(
                  child: Obx(
            () => ListView.builder(
              itemBuilder: (BuildContext context, int index) {
                return GestureDetector(
                  onTap: () {
                    Get.toNamed('/community?id=' + communities.value[index]['id']);
                  },
                  child: mine.Card(cur: Community.fromJson(communities.value[index])),
                );
              },
              itemCount: communities.value.length,
              itemExtent: 90,
              shrinkWrap: true,
            ),
          ))),
        ],
      ),
    );
  }
}

class Backed extends StatefulWidget {
  const Backed({Key? key}) : super(key: key);

  @override
  State<Backed> createState() => _BackedState();
}

class _BackedState extends State<Backed> {
  @override
  Widget build(BuildContext context) {
    return const Text("This is backed");
  }
}

class About extends StatelessWidget {
  const About({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const Text("This is about");
  }
}
