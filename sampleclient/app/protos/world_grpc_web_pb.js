/**
 * @fileoverview gRPC-Web generated client stub for protos
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var position_pb = require('./position_pb.js')
const proto = {};
proto.protos = require('./world_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.protos.WorldListenClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.protos.WorldListenPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.protos.WorldRequest,
 *   !proto.protos.World>}
 */
const methodInfo_WorldListen_WorldStartListen = new grpc.web.AbstractClientBase.MethodInfo(
  proto.protos.World,
  /** @param {!proto.protos.WorldRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.World.deserializeBinary
);


/**
 * @param {!proto.protos.WorldRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.protos.World>}
 *     The XHR Node Readable Stream
 */
proto.protos.WorldListenClient.prototype.worldStartListen =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/protos.WorldListen/WorldStartListen',
      request,
      metadata || {},
      methodInfo_WorldListen_WorldStartListen);
};


/**
 * @param {!proto.protos.WorldRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.protos.World>}
 *     The XHR Node Readable Stream
 */
proto.protos.WorldListenPromiseClient.prototype.worldStartListen =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/protos.WorldListen/WorldStartListen',
      request,
      metadata || {},
      methodInfo_WorldListen_WorldStartListen);
};


module.exports = proto.protos;

