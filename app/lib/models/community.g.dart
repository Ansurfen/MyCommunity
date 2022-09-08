// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'community.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Community _$CommunityFromJson(Map<String, dynamic> json) => Community(
      name: json['name'] as String?,
      id: json['id'] as String?,
      image: json['image'] as String?,
      context: json['context'] as String?,
      admins: json['admins'] as String?,
      hostname: json['hostname'] as String?,
      members: json['members'] as String?,
      notes: json['notes'] as String?,
      posts: json['posts'] as String?,
      status: json['status'] as int?,
      timestamp: json['timestamp'] as String?,
      tags: json['tags'] as String?,
    );

Map<String, dynamic> _$CommunityToJson(Community instance) => <String, dynamic>{
      'status': instance.status,
      'id': instance.id,
      'context': instance.context,
      'name': instance.name,
      'image': instance.image,
      'timestamp': instance.timestamp,
      'hostname': instance.hostname,
      'admins': instance.admins,
      'members': instance.members,
      'notes': instance.notes,
      'posts': instance.posts,
      'tags': instance.tags,
    };
