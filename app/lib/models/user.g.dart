// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'user.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

User _$UserFromJson(Map<String, dynamic> json) => User(
      username: json['username'] as String?,
      password: json['password'] as String?,
    );

Map<String, dynamic> _$UserToJson(User instance) => <String, dynamic>{
      'username': instance.username,
      'password': instance.password,
    };

UserInfo _$UserInfoFromJson(Map<String, dynamic> json) => UserInfo(
      right: json['right'] as int?,
      id: json['id'] as int?,
      username: json['username'] as String?,
      alias: json['alias'] as String?,
      telephone: json['telephone'] as String?,
      email: json['email'] as String?,
      school: json['school'] as String?,
      profile: json['profile'] as int?,
    );

Map<String, dynamic> _$UserInfoToJson(UserInfo instance) => <String, dynamic>{
      'right': instance.right,
      'id': instance.id,
      'username': instance.username,
      'alias': instance.alias,
      'telephone': instance.telephone,
      'email': instance.email,
      'school': instance.school,
      'profile': instance.profile,
    };
