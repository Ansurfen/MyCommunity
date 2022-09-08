import 'package:json_annotation/json_annotation.dart';

part 'community.g.dart';

@JsonSerializable()
class Community {
  @JsonKey(name: "status")
  int? status;
  @JsonKey(name: "id")
  String? id;
  @JsonKey(name: "context")
  String? context;
  @JsonKey(name: "name")
  String? name;
  @JsonKey(name: "image")
  String? image;
  @JsonKey(name: "timestamp")
  String? timestamp;
  @JsonKey(name: "hostname")
  String? hostname;
  @JsonKey(name: "admins")
  String? admins;
  @JsonKey(name: "members")
  String? members;
  @JsonKey(name: "notes")
  String? notes;
  @JsonKey(name: "posts")
  String? posts;
  @JsonKey(name: "tags")
  String? tags;

  Community(
      {this.name,
      this.id,
      this.image,
      this.context,
      this.admins,
      this.hostname,
      this.members,
      this.notes,
      this.posts,
      this.status,
      this.timestamp,
      this.tags});

  factory Community.fromJson(Map<String, dynamic> json) =>
      _$CommunityFromJson(json);

  Map<String, dynamic> toJson() => _$CommunityToJson(this);
}

class Applications {
  int? type;
  int? status;
  String? first;
  String? second;
  String? context;
}

class Post {
  int? score;
  String? id;
  String? title;
  String? author;
  String? timestamp;
  String? context;
  List<String>? tags;
}

class Note {
  String? title;
  String? context;
  String? timestamp;
}
