import 'package:json_annotation/json_annotation.dart';

part 'user.g.dart';

@JsonSerializable()
class User {
  @JsonKey(name: 'username')
  String? username;
  @JsonKey(name: 'password')
  String? password;

  User({this.username, this.password});

  factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);

  Map<String, dynamic> toJson() => _$UserToJson(this);
}

@JsonSerializable()
class UserInfo {
  @JsonKey(name: "right")
  int? right;
  @JsonKey(name: "id")
  int? id;
  @JsonKey(name: "username")
  String? username;
  @JsonKey(name: "alias")
  String? alias;
  @JsonKey(name: "telephone")
  String? telephone;
  @JsonKey(name: "email")
  String? email;
  @JsonKey(name: "school")
  String? school;
  @JsonKey(name: "profile")
  int? profile;

  UserInfo(
      {this.right,
      this.id,
      this.username,
      this.alias,
      this.telephone,
      this.email,
      this.school,
      this.profile});

  factory UserInfo.fromJson(Map<String, dynamic> json) => _$UserInfoFromJson(json);

  Map<String, dynamic> toJson() => _$UserInfoToJson(this);
}
