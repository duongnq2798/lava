/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var gogoproto_gogo_pb = require('../../../gogoproto/gogo_pb.js');
var lavanet_lava_plans_policy_pb = require('../../../lavanet/lava/plans/policy_pb.js');
goog.exportSymbol('proto.lavanet.lava.projects.Project', null, global);
goog.exportSymbol('proto.lavanet.lava.projects.ProjectData', null, global);
goog.exportSymbol('proto.lavanet.lava.projects.ProjectKey', null, global);
goog.exportSymbol('proto.lavanet.lava.projects.ProjectKey.Type', null, global);
goog.exportSymbol('proto.lavanet.lava.projects.ProtoDeveloperData', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.lavanet.lava.projects.Project = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.lavanet.lava.projects.Project.repeatedFields_, null);
};
goog.inherits(proto.lavanet.lava.projects.Project, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.lavanet.lava.projects.Project.displayName = 'proto.lavanet.lava.projects.Project';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.lavanet.lava.projects.Project.repeatedFields_ = [5];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.lavanet.lava.projects.Project.prototype.toObject = function(opt_includeInstance) {
  return proto.lavanet.lava.projects.Project.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.lavanet.lava.projects.Project} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.projects.Project.toObject = function(includeInstance, msg) {
  var f, obj = {
    index: jspb.Message.getFieldWithDefault(msg, 1, ""),
    subscription: jspb.Message.getFieldWithDefault(msg, 2, ""),
    enabled: jspb.Message.getFieldWithDefault(msg, 4, false),
    projectKeysList: jspb.Message.toObjectList(msg.getProjectKeysList(),
    proto.lavanet.lava.projects.ProjectKey.toObject, includeInstance),
    adminPolicy: (f = msg.getAdminPolicy()) && lavanet_lava_plans_policy_pb.Policy.toObject(includeInstance, f),
    usedCu: jspb.Message.getFieldWithDefault(msg, 7, 0),
    subscriptionPolicy: (f = msg.getSubscriptionPolicy()) && lavanet_lava_plans_policy_pb.Policy.toObject(includeInstance, f),
    snapshot: jspb.Message.getFieldWithDefault(msg, 9, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.lavanet.lava.projects.Project}
 */
proto.lavanet.lava.projects.Project.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.lavanet.lava.projects.Project;
  return proto.lavanet.lava.projects.Project.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.lavanet.lava.projects.Project} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.lavanet.lava.projects.Project}
 */
proto.lavanet.lava.projects.Project.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setIndex(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setSubscription(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setEnabled(value);
      break;
    case 5:
      var value = new proto.lavanet.lava.projects.ProjectKey;
      reader.readMessage(value,proto.lavanet.lava.projects.ProjectKey.deserializeBinaryFromReader);
      msg.addProjectKeys(value);
      break;
    case 6:
      var value = new lavanet_lava_plans_policy_pb.Policy;
      reader.readMessage(value,lavanet_lava_plans_policy_pb.Policy.deserializeBinaryFromReader);
      msg.setAdminPolicy(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setUsedCu(value);
      break;
    case 8:
      var value = new lavanet_lava_plans_policy_pb.Policy;
      reader.readMessage(value,lavanet_lava_plans_policy_pb.Policy.deserializeBinaryFromReader);
      msg.setSubscriptionPolicy(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setSnapshot(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.lavanet.lava.projects.Project.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.lavanet.lava.projects.Project.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.lavanet.lava.projects.Project} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.projects.Project.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getIndex();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getSubscription();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getEnabled();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
  f = message.getProjectKeysList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      5,
      f,
      proto.lavanet.lava.projects.ProjectKey.serializeBinaryToWriter
    );
  }
  f = message.getAdminPolicy();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      lavanet_lava_plans_policy_pb.Policy.serializeBinaryToWriter
    );
  }
  f = message.getUsedCu();
  if (f !== 0) {
    writer.writeUint64(
      7,
      f
    );
  }
  f = message.getSubscriptionPolicy();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      lavanet_lava_plans_policy_pb.Policy.serializeBinaryToWriter
    );
  }
  f = message.getSnapshot();
  if (f !== 0) {
    writer.writeUint64(
      9,
      f
    );
  }
};


/**
 * optional string index = 1;
 * @return {string}
 */
proto.lavanet.lava.projects.Project.prototype.getIndex = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.lavanet.lava.projects.Project.prototype.setIndex = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string subscription = 2;
 * @return {string}
 */
proto.lavanet.lava.projects.Project.prototype.getSubscription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.lavanet.lava.projects.Project.prototype.setSubscription = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional bool enabled = 4;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.lavanet.lava.projects.Project.prototype.getEnabled = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 4, false));
};


/** @param {boolean} value */
proto.lavanet.lava.projects.Project.prototype.setEnabled = function(value) {
  jspb.Message.setProto3BooleanField(this, 4, value);
};


/**
 * repeated ProjectKey project_keys = 5;
 * @return {!Array<!proto.lavanet.lava.projects.ProjectKey>}
 */
proto.lavanet.lava.projects.Project.prototype.getProjectKeysList = function() {
  return /** @type{!Array<!proto.lavanet.lava.projects.ProjectKey>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.lavanet.lava.projects.ProjectKey, 5));
};


/** @param {!Array<!proto.lavanet.lava.projects.ProjectKey>} value */
proto.lavanet.lava.projects.Project.prototype.setProjectKeysList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 5, value);
};


/**
 * @param {!proto.lavanet.lava.projects.ProjectKey=} opt_value
 * @param {number=} opt_index
 * @return {!proto.lavanet.lava.projects.ProjectKey}
 */
proto.lavanet.lava.projects.Project.prototype.addProjectKeys = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 5, opt_value, proto.lavanet.lava.projects.ProjectKey, opt_index);
};


proto.lavanet.lava.projects.Project.prototype.clearProjectKeysList = function() {
  this.setProjectKeysList([]);
};


/**
 * optional lavanet.lava.plans.Policy admin_policy = 6;
 * @return {?proto.lavanet.lava.plans.Policy}
 */
proto.lavanet.lava.projects.Project.prototype.getAdminPolicy = function() {
  return /** @type{?proto.lavanet.lava.plans.Policy} */ (
    jspb.Message.getWrapperField(this, lavanet_lava_plans_policy_pb.Policy, 6));
};


/** @param {?proto.lavanet.lava.plans.Policy|undefined} value */
proto.lavanet.lava.projects.Project.prototype.setAdminPolicy = function(value) {
  jspb.Message.setWrapperField(this, 6, value);
};


proto.lavanet.lava.projects.Project.prototype.clearAdminPolicy = function() {
  this.setAdminPolicy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.lavanet.lava.projects.Project.prototype.hasAdminPolicy = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional uint64 used_cu = 7;
 * @return {number}
 */
proto.lavanet.lava.projects.Project.prototype.getUsedCu = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/** @param {number} value */
proto.lavanet.lava.projects.Project.prototype.setUsedCu = function(value) {
  jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * optional lavanet.lava.plans.Policy subscription_policy = 8;
 * @return {?proto.lavanet.lava.plans.Policy}
 */
proto.lavanet.lava.projects.Project.prototype.getSubscriptionPolicy = function() {
  return /** @type{?proto.lavanet.lava.plans.Policy} */ (
    jspb.Message.getWrapperField(this, lavanet_lava_plans_policy_pb.Policy, 8));
};


/** @param {?proto.lavanet.lava.plans.Policy|undefined} value */
proto.lavanet.lava.projects.Project.prototype.setSubscriptionPolicy = function(value) {
  jspb.Message.setWrapperField(this, 8, value);
};


proto.lavanet.lava.projects.Project.prototype.clearSubscriptionPolicy = function() {
  this.setSubscriptionPolicy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.lavanet.lava.projects.Project.prototype.hasSubscriptionPolicy = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional uint64 snapshot = 9;
 * @return {number}
 */
proto.lavanet.lava.projects.Project.prototype.getSnapshot = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/** @param {number} value */
proto.lavanet.lava.projects.Project.prototype.setSnapshot = function(value) {
  jspb.Message.setProto3IntField(this, 9, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.lavanet.lava.projects.ProjectKey = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.lavanet.lava.projects.ProjectKey, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.lavanet.lava.projects.ProjectKey.displayName = 'proto.lavanet.lava.projects.ProjectKey';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.lavanet.lava.projects.ProjectKey.prototype.toObject = function(opt_includeInstance) {
  return proto.lavanet.lava.projects.ProjectKey.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.lavanet.lava.projects.ProjectKey} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.projects.ProjectKey.toObject = function(includeInstance, msg) {
  var f, obj = {
    key: jspb.Message.getFieldWithDefault(msg, 1, ""),
    kinds: jspb.Message.getFieldWithDefault(msg, 4, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.lavanet.lava.projects.ProjectKey}
 */
proto.lavanet.lava.projects.ProjectKey.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.lavanet.lava.projects.ProjectKey;
  return proto.lavanet.lava.projects.ProjectKey.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.lavanet.lava.projects.ProjectKey} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.lavanet.lava.projects.ProjectKey}
 */
proto.lavanet.lava.projects.ProjectKey.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setKey(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setKinds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.lavanet.lava.projects.ProjectKey.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.lavanet.lava.projects.ProjectKey.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.lavanet.lava.projects.ProjectKey} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.projects.ProjectKey.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getKey();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getKinds();
  if (f !== 0) {
    writer.writeUint32(
      4,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.lavanet.lava.projects.ProjectKey.Type = {
  NONE: 0,
  ADMIN: 1,
  DEVELOPER: 2
};

/**
 * optional string key = 1;
 * @return {string}
 */
proto.lavanet.lava.projects.ProjectKey.prototype.getKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.lavanet.lava.projects.ProjectKey.prototype.setKey = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional uint32 kinds = 4;
 * @return {number}
 */
proto.lavanet.lava.projects.ProjectKey.prototype.getKinds = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.lavanet.lava.projects.ProjectKey.prototype.setKinds = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.lavanet.lava.projects.ProtoDeveloperData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.lavanet.lava.projects.ProtoDeveloperData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.lavanet.lava.projects.ProtoDeveloperData.displayName = 'proto.lavanet.lava.projects.ProtoDeveloperData';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.lavanet.lava.projects.ProtoDeveloperData.prototype.toObject = function(opt_includeInstance) {
  return proto.lavanet.lava.projects.ProtoDeveloperData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.lavanet.lava.projects.ProtoDeveloperData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.projects.ProtoDeveloperData.toObject = function(includeInstance, msg) {
  var f, obj = {
    projectid: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.lavanet.lava.projects.ProtoDeveloperData}
 */
proto.lavanet.lava.projects.ProtoDeveloperData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.lavanet.lava.projects.ProtoDeveloperData;
  return proto.lavanet.lava.projects.ProtoDeveloperData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.lavanet.lava.projects.ProtoDeveloperData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.lavanet.lava.projects.ProtoDeveloperData}
 */
proto.lavanet.lava.projects.ProtoDeveloperData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProjectid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.lavanet.lava.projects.ProtoDeveloperData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.lavanet.lava.projects.ProtoDeveloperData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.lavanet.lava.projects.ProtoDeveloperData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.projects.ProtoDeveloperData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProjectid();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string projectID = 1;
 * @return {string}
 */
proto.lavanet.lava.projects.ProtoDeveloperData.prototype.getProjectid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.lavanet.lava.projects.ProtoDeveloperData.prototype.setProjectid = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.lavanet.lava.projects.ProjectData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.lavanet.lava.projects.ProjectData.repeatedFields_, null);
};
goog.inherits(proto.lavanet.lava.projects.ProjectData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.lavanet.lava.projects.ProjectData.displayName = 'proto.lavanet.lava.projects.ProjectData';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.lavanet.lava.projects.ProjectData.repeatedFields_ = [4];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.lavanet.lava.projects.ProjectData.prototype.toObject = function(opt_includeInstance) {
  return proto.lavanet.lava.projects.ProjectData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.lavanet.lava.projects.ProjectData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.projects.ProjectData.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    enabled: jspb.Message.getFieldWithDefault(msg, 3, false),
    projectkeysList: jspb.Message.toObjectList(msg.getProjectkeysList(),
    proto.lavanet.lava.projects.ProjectKey.toObject, includeInstance),
    policy: (f = msg.getPolicy()) && lavanet_lava_plans_policy_pb.Policy.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.lavanet.lava.projects.ProjectData}
 */
proto.lavanet.lava.projects.ProjectData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.lavanet.lava.projects.ProjectData;
  return proto.lavanet.lava.projects.ProjectData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.lavanet.lava.projects.ProjectData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.lavanet.lava.projects.ProjectData}
 */
proto.lavanet.lava.projects.ProjectData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setEnabled(value);
      break;
    case 4:
      var value = new proto.lavanet.lava.projects.ProjectKey;
      reader.readMessage(value,proto.lavanet.lava.projects.ProjectKey.deserializeBinaryFromReader);
      msg.addProjectkeys(value);
      break;
    case 5:
      var value = new lavanet_lava_plans_policy_pb.Policy;
      reader.readMessage(value,lavanet_lava_plans_policy_pb.Policy.deserializeBinaryFromReader);
      msg.setPolicy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.lavanet.lava.projects.ProjectData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.lavanet.lava.projects.ProjectData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.lavanet.lava.projects.ProjectData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.projects.ProjectData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getEnabled();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
  f = message.getProjectkeysList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      4,
      f,
      proto.lavanet.lava.projects.ProjectKey.serializeBinaryToWriter
    );
  }
  f = message.getPolicy();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      lavanet_lava_plans_policy_pb.Policy.serializeBinaryToWriter
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.lavanet.lava.projects.ProjectData.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.lavanet.lava.projects.ProjectData.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional bool enabled = 3;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.lavanet.lava.projects.ProjectData.prototype.getEnabled = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 3, false));
};


/** @param {boolean} value */
proto.lavanet.lava.projects.ProjectData.prototype.setEnabled = function(value) {
  jspb.Message.setProto3BooleanField(this, 3, value);
};


/**
 * repeated ProjectKey projectKeys = 4;
 * @return {!Array<!proto.lavanet.lava.projects.ProjectKey>}
 */
proto.lavanet.lava.projects.ProjectData.prototype.getProjectkeysList = function() {
  return /** @type{!Array<!proto.lavanet.lava.projects.ProjectKey>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.lavanet.lava.projects.ProjectKey, 4));
};


/** @param {!Array<!proto.lavanet.lava.projects.ProjectKey>} value */
proto.lavanet.lava.projects.ProjectData.prototype.setProjectkeysList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 4, value);
};


/**
 * @param {!proto.lavanet.lava.projects.ProjectKey=} opt_value
 * @param {number=} opt_index
 * @return {!proto.lavanet.lava.projects.ProjectKey}
 */
proto.lavanet.lava.projects.ProjectData.prototype.addProjectkeys = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 4, opt_value, proto.lavanet.lava.projects.ProjectKey, opt_index);
};


proto.lavanet.lava.projects.ProjectData.prototype.clearProjectkeysList = function() {
  this.setProjectkeysList([]);
};


/**
 * optional lavanet.lava.plans.Policy policy = 5;
 * @return {?proto.lavanet.lava.plans.Policy}
 */
proto.lavanet.lava.projects.ProjectData.prototype.getPolicy = function() {
  return /** @type{?proto.lavanet.lava.plans.Policy} */ (
    jspb.Message.getWrapperField(this, lavanet_lava_plans_policy_pb.Policy, 5));
};


/** @param {?proto.lavanet.lava.plans.Policy|undefined} value */
proto.lavanet.lava.projects.ProjectData.prototype.setPolicy = function(value) {
  jspb.Message.setWrapperField(this, 5, value);
};


proto.lavanet.lava.projects.ProjectData.prototype.clearPolicy = function() {
  this.setPolicy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.lavanet.lava.projects.ProjectData.prototype.hasPolicy = function() {
  return jspb.Message.getField(this, 5) != null;
};


goog.object.extend(exports, proto.lavanet.lava.projects);