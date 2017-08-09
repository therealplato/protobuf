/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, {
/******/ 				configurable: false,
/******/ 				enumerable: true,
/******/ 				get: getter
/******/ 			});
/******/ 		}
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = 3);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ (function(module, exports) {

module.exports = window;

/***/ }),
/* 1 */
/***/ (function(module, exports, __webpack_require__) {

/**
 * @fileoverview
 * @enhanceable
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = __webpack_require__(0);
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.multitest.Multi3', null, global);
goog.exportSymbol('proto.multitest.Multi3.HatType', null, global);

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
proto.multitest.Multi3 = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.multitest.Multi3, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.multitest.Multi3.displayName = 'proto.multitest.Multi3';
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
proto.multitest.Multi3.prototype.toObject = function(opt_includeInstance) {
  return proto.multitest.Multi3.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.multitest.Multi3} msg The msg instance to transform.
 * @return {!Object}
 */
proto.multitest.Multi3.toObject = function(includeInstance, msg) {
  var f, obj = {
    hatType: jspb.Message.getFieldWithDefault(msg, 1, 0)
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
 * @return {!proto.multitest.Multi3}
 */
proto.multitest.Multi3.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.multitest.Multi3;
  return proto.multitest.Multi3.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.multitest.Multi3} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.multitest.Multi3}
 */
proto.multitest.Multi3.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.multitest.Multi3.HatType} */ (reader.readEnum());
      msg.setHatType(value);
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
proto.multitest.Multi3.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.multitest.Multi3.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.multitest.Multi3} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.multitest.Multi3.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHatType();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.multitest.Multi3.HatType = {
  FEDORA: 0,
  FEZ: 1
};

/**
 * optional HatType hat_type = 1;
 * @return {!proto.multitest.Multi3.HatType}
 */
proto.multitest.Multi3.prototype.getHatType = function() {
  return /** @type {!proto.multitest.Multi3.HatType} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.multitest.Multi3.HatType} value */
proto.multitest.Multi3.prototype.setHatType = function(value) {
  jspb.Message.setField(this, 1, value);
};


goog.object.extend(exports, proto.multitest);


/***/ }),
/* 2 */
/***/ (function(module, exports, __webpack_require__) {

/**
 * @fileoverview
 * @enhanceable
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = __webpack_require__(0);
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.multitest.Multi2', null, global);
goog.exportSymbol('proto.multitest.Multi2.Color', null, global);

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
proto.multitest.Multi2 = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.multitest.Multi2, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.multitest.Multi2.displayName = 'proto.multitest.Multi2';
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
proto.multitest.Multi2.prototype.toObject = function(opt_includeInstance) {
  return proto.multitest.Multi2.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.multitest.Multi2} msg The msg instance to transform.
 * @return {!Object}
 */
proto.multitest.Multi2.toObject = function(includeInstance, msg) {
  var f, obj = {
    requiredValue: jspb.Message.getFieldWithDefault(msg, 1, 0),
    color: jspb.Message.getFieldWithDefault(msg, 2, 0)
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
 * @return {!proto.multitest.Multi2}
 */
proto.multitest.Multi2.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.multitest.Multi2;
  return proto.multitest.Multi2.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.multitest.Multi2} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.multitest.Multi2}
 */
proto.multitest.Multi2.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRequiredValue(value);
      break;
    case 2:
      var value = /** @type {!proto.multitest.Multi2.Color} */ (reader.readEnum());
      msg.setColor(value);
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
proto.multitest.Multi2.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.multitest.Multi2.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.multitest.Multi2} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.multitest.Multi2.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRequiredValue();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getColor();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.multitest.Multi2.Color = {
  BLUE: 0,
  GREEN: 1,
  RED: 2
};

/**
 * optional int32 required_value = 1;
 * @return {number}
 */
proto.multitest.Multi2.prototype.getRequiredValue = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.multitest.Multi2.prototype.setRequiredValue = function(value) {
  jspb.Message.setField(this, 1, value);
};


/**
 * optional Color color = 2;
 * @return {!proto.multitest.Multi2.Color}
 */
proto.multitest.Multi2.prototype.getColor = function() {
  return /** @type {!proto.multitest.Multi2.Color} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {!proto.multitest.Multi2.Color} value */
proto.multitest.Multi2.prototype.setColor = function(value) {
  jspb.Message.setField(this, 2, value);
};


goog.object.extend(exports, proto.multitest);


/***/ }),
/* 3 */
/***/ (function(module, exports, __webpack_require__) {

__webpack_require__(1);
__webpack_require__(2);
module.exports = __webpack_require__(4);


/***/ }),
/* 4 */
/***/ (function(module, exports, __webpack_require__) {

/**
 * @fileoverview
 * @enhanceable
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = __webpack_require__(0);
var goog = jspb;
var global = Function('return this')();

var multi_multi2_pb = __webpack_require__(2);
var multi_multi3_pb = __webpack_require__(1);
goog.exportSymbol('proto.multitest.Multi1', null, global);

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
proto.multitest.Multi1 = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.multitest.Multi1, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.multitest.Multi1.displayName = 'proto.multitest.Multi1';
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
proto.multitest.Multi1.prototype.toObject = function(opt_includeInstance) {
  return proto.multitest.Multi1.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.multitest.Multi1} msg The msg instance to transform.
 * @return {!Object}
 */
proto.multitest.Multi1.toObject = function(includeInstance, msg) {
  var f, obj = {
    multi2: (f = msg.getMulti2()) && multi_multi2_pb.Multi2.toObject(includeInstance, f),
    color: jspb.Message.getFieldWithDefault(msg, 2, 0),
    hatType: jspb.Message.getFieldWithDefault(msg, 3, 0)
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
 * @return {!proto.multitest.Multi1}
 */
proto.multitest.Multi1.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.multitest.Multi1;
  return proto.multitest.Multi1.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.multitest.Multi1} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.multitest.Multi1}
 */
proto.multitest.Multi1.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new multi_multi2_pb.Multi2;
      reader.readMessage(value,multi_multi2_pb.Multi2.deserializeBinaryFromReader);
      msg.setMulti2(value);
      break;
    case 2:
      var value = /** @type {!proto.multitest.Multi2.Color} */ (reader.readEnum());
      msg.setColor(value);
      break;
    case 3:
      var value = /** @type {!proto.multitest.Multi3.HatType} */ (reader.readEnum());
      msg.setHatType(value);
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
proto.multitest.Multi1.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.multitest.Multi1.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.multitest.Multi1} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.multitest.Multi1.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMulti2();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      multi_multi2_pb.Multi2.serializeBinaryToWriter
    );
  }
  f = message.getColor();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getHatType();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
};


/**
 * optional Multi2 multi2 = 1;
 * @return {?proto.multitest.Multi2}
 */
proto.multitest.Multi1.prototype.getMulti2 = function() {
  return /** @type{?proto.multitest.Multi2} */ (
    jspb.Message.getWrapperField(this, multi_multi2_pb.Multi2, 1));
};


/** @param {?proto.multitest.Multi2|undefined} value */
proto.multitest.Multi1.prototype.setMulti2 = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.multitest.Multi1.prototype.clearMulti2 = function() {
  this.setMulti2(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.multitest.Multi1.prototype.hasMulti2 = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional Multi2.Color color = 2;
 * @return {!proto.multitest.Multi2.Color}
 */
proto.multitest.Multi1.prototype.getColor = function() {
  return /** @type {!proto.multitest.Multi2.Color} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {!proto.multitest.Multi2.Color} value */
proto.multitest.Multi1.prototype.setColor = function(value) {
  jspb.Message.setField(this, 2, value);
};


/**
 * optional Multi3.HatType hat_type = 3;
 * @return {!proto.multitest.Multi3.HatType}
 */
proto.multitest.Multi1.prototype.getHatType = function() {
  return /** @type {!proto.multitest.Multi3.HatType} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {!proto.multitest.Multi3.HatType} value */
proto.multitest.Multi1.prototype.setHatType = function(value) {
  jspb.Message.setField(this, 3, value);
};


goog.object.extend(exports, proto.multitest);


/***/ })
/******/ ]);