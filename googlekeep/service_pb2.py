# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: service.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='service.proto',
  package='',
  syntax='proto3',
  serialized_options=_b('Z3github.com/luckyshmo/api-example/api/keepServiceApi'),
  serialized_pb=_b('\n\rservice.proto\";\n\rSearchRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\x12\r\n\x05token\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\"\x1c\n\x0cSearchResult\x12\x0c\n\x04word\x18\x03 \x03(\t23\n\x04Keep\x12+\n\x08GetWords\x12\x0e.SearchRequest\x1a\r.SearchResult\"\x00\x42\x35Z3github.com/luckyshmo/api-example/api/keepServiceApib\x06proto3')
)




_SEARCHREQUEST = _descriptor.Descriptor(
  name='SearchRequest',
  full_name='SearchRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='email', full_name='SearchRequest.email', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='token', full_name='SearchRequest.token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='SearchRequest.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=17,
  serialized_end=76,
)


_SEARCHRESULT = _descriptor.Descriptor(
  name='SearchResult',
  full_name='SearchResult',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='word', full_name='SearchResult.word', index=0,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=78,
  serialized_end=106,
)

DESCRIPTOR.message_types_by_name['SearchRequest'] = _SEARCHREQUEST
DESCRIPTOR.message_types_by_name['SearchResult'] = _SEARCHRESULT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SearchRequest = _reflection.GeneratedProtocolMessageType('SearchRequest', (_message.Message,), dict(
  DESCRIPTOR = _SEARCHREQUEST,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:SearchRequest)
  ))
_sym_db.RegisterMessage(SearchRequest)

SearchResult = _reflection.GeneratedProtocolMessageType('SearchResult', (_message.Message,), dict(
  DESCRIPTOR = _SEARCHRESULT,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:SearchResult)
  ))
_sym_db.RegisterMessage(SearchResult)


DESCRIPTOR._options = None

_KEEP = _descriptor.ServiceDescriptor(
  name='Keep',
  full_name='Keep',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=108,
  serialized_end=159,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetWords',
    full_name='Keep.GetWords',
    index=0,
    containing_service=None,
    input_type=_SEARCHREQUEST,
    output_type=_SEARCHRESULT,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_KEEP)

DESCRIPTOR.services_by_name['Keep'] = _KEEP

# @@protoc_insertion_point(module_scope)
