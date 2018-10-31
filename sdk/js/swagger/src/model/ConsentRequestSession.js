/**
 * ORY Hydra - Cloud Native OAuth 2.0 and OpenID Connect Server
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here. Keep in mind that this document reflects the latest branch, always. Support for versioned documentation is coming in the future.
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.OryHydraCloudNativeOAuth20AndOpenIdConnectServer) {
      root.OryHydraCloudNativeOAuth20AndOpenIdConnectServer = {};
    }
    root.OryHydraCloudNativeOAuth20AndOpenIdConnectServer.ConsentRequestSession = factory(root.OryHydraCloudNativeOAuth20AndOpenIdConnectServer.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The ConsentRequestSession model module.
   * @module model/ConsentRequestSession
   * @version Latest
   */

  /**
   * Constructs a new <code>ConsentRequestSession</code>.
   * @alias module:model/ConsentRequestSession
   * @class
   */
  var exports = function() {
    var _this = this;



  };

  /**
   * Constructs a <code>ConsentRequestSession</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ConsentRequestSession} obj Optional instance to populate.
   * @return {module:model/ConsentRequestSession} The populated <code>ConsentRequestSession</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('access_token')) {
        obj['access_token'] = ApiClient.convertToType(data['access_token'], {'String': Object});
      }
      if (data.hasOwnProperty('id_token')) {
        obj['id_token'] = ApiClient.convertToType(data['id_token'], {'String': Object});
      }
    }
    return obj;
  }

  /**
   * AccessToken sets session data for the access and refresh token, as well as any future tokens issued by the refresh grant. Keep in mind that this data will be available to anyone performing OAuth 2.0 Challenge Introspection. If only your services can perform OAuth 2.0 Challenge Introspection, this is usually fine. But if third parties can access that endpoint as well, sensitive data from the session might be exposed to them. Use with care!
   * @member {Object.<String, Object>} access_token
   */
  exports.prototype['access_token'] = undefined;
  /**
   * IDToken sets session data for the OpenID Connect ID token. Keep in mind that the session'id payloads are readable by anyone that has access to the ID Challenge. Use with care!
   * @member {Object.<String, Object>} id_token
   */
  exports.prototype['id_token'] = undefined;



  return exports;
}));


