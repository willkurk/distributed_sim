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
proto.protos = require('./world_render_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.protos.WorldSyncRenderClient =
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
proto.protos.WorldSyncRenderPromiseClient =
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
 *   !proto.protos.WorldRenderRequest,
 *   !proto.protos.WorldRender>}
 */
const methodInfo_WorldSyncRender_WorldStartRender = new grpc.web.AbstractClientBase.MethodInfo(
  proto.protos.WorldRender,
  /** @param {!proto.protos.WorldRenderRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.WorldRender.deserializeBinary
);


/**
 * @param {!proto.protos.WorldRenderRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.protos.WorldRender>}
 *     The XHR Node Readable Stream
 */
proto.protos.WorldSyncRenderClient.prototype.worldStartRender =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/protos.WorldSyncRender/WorldStartRender',
      request,
      metadata || {},
      methodInfo_WorldSyncRender_WorldStartRender);
};


/**
 * @param {!proto.protos.WorldRenderRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.protos.WorldRender>}
 *     The XHR Node Readable Stream
 */
proto.protos.WorldSyncRenderPromiseClient.prototype.worldStartRender =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/protos.WorldSyncRender/WorldStartRender',
      request,
      metadata || {},
      methodInfo_WorldSyncRender_WorldStartRender);
};


module.exports = proto.protos;

