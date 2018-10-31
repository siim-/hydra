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
    define(['ApiClient', 'model/JSONWebKeySet'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./JSONWebKeySet'));
  } else {
    // Browser globals (root is window)
    if (!root.OryHydraCloudNativeOAuth20AndOpenIdConnectServer) {
      root.OryHydraCloudNativeOAuth20AndOpenIdConnectServer = {};
    }
    root.OryHydraCloudNativeOAuth20AndOpenIdConnectServer.OAuth2Client = factory(root.OryHydraCloudNativeOAuth20AndOpenIdConnectServer.ApiClient, root.OryHydraCloudNativeOAuth20AndOpenIdConnectServer.JSONWebKeySet);
  }
}(this, function(ApiClient, JSONWebKeySet) {
  'use strict';




  /**
   * The OAuth2Client model module.
   * @module model/OAuth2Client
   * @version Latest
   */

  /**
   * Constructs a new <code>OAuth2Client</code>.
   * @alias module:model/OAuth2Client
   * @class
   */
  var exports = function() {
    var _this = this;

























  };

  /**
   * Constructs a <code>OAuth2Client</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/OAuth2Client} obj Optional instance to populate.
   * @return {module:model/OAuth2Client} The populated <code>OAuth2Client</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('allowed_cors_origins')) {
        obj['allowed_cors_origins'] = ApiClient.convertToType(data['allowed_cors_origins'], ['String']);
      }
      if (data.hasOwnProperty('audience')) {
        obj['audience'] = ApiClient.convertToType(data['audience'], ['String']);
      }
      if (data.hasOwnProperty('client_id')) {
        obj['client_id'] = ApiClient.convertToType(data['client_id'], 'String');
      }
      if (data.hasOwnProperty('client_name')) {
        obj['client_name'] = ApiClient.convertToType(data['client_name'], 'String');
      }
      if (data.hasOwnProperty('client_secret')) {
        obj['client_secret'] = ApiClient.convertToType(data['client_secret'], 'String');
      }
      if (data.hasOwnProperty('client_secret_expires_at')) {
        obj['client_secret_expires_at'] = ApiClient.convertToType(data['client_secret_expires_at'], 'Number');
      }
      if (data.hasOwnProperty('client_uri')) {
        obj['client_uri'] = ApiClient.convertToType(data['client_uri'], 'String');
      }
      if (data.hasOwnProperty('contacts')) {
        obj['contacts'] = ApiClient.convertToType(data['contacts'], ['String']);
      }
      if (data.hasOwnProperty('grant_types')) {
        obj['grant_types'] = ApiClient.convertToType(data['grant_types'], ['String']);
      }
      if (data.hasOwnProperty('jwks')) {
        obj['jwks'] = JSONWebKeySet.constructFromObject(data['jwks']);
      }
      if (data.hasOwnProperty('jwks_uri')) {
        obj['jwks_uri'] = ApiClient.convertToType(data['jwks_uri'], 'String');
      }
      if (data.hasOwnProperty('logo_uri')) {
        obj['logo_uri'] = ApiClient.convertToType(data['logo_uri'], 'String');
      }
      if (data.hasOwnProperty('owner')) {
        obj['owner'] = ApiClient.convertToType(data['owner'], 'String');
      }
      if (data.hasOwnProperty('policy_uri')) {
        obj['policy_uri'] = ApiClient.convertToType(data['policy_uri'], 'String');
      }
      if (data.hasOwnProperty('redirect_uris')) {
        obj['redirect_uris'] = ApiClient.convertToType(data['redirect_uris'], ['String']);
      }
      if (data.hasOwnProperty('request_object_signing_alg')) {
        obj['request_object_signing_alg'] = ApiClient.convertToType(data['request_object_signing_alg'], 'String');
      }
      if (data.hasOwnProperty('request_uris')) {
        obj['request_uris'] = ApiClient.convertToType(data['request_uris'], ['String']);
      }
      if (data.hasOwnProperty('response_types')) {
        obj['response_types'] = ApiClient.convertToType(data['response_types'], ['String']);
      }
      if (data.hasOwnProperty('scope')) {
        obj['scope'] = ApiClient.convertToType(data['scope'], 'String');
      }
      if (data.hasOwnProperty('sector_identifier_uri')) {
        obj['sector_identifier_uri'] = ApiClient.convertToType(data['sector_identifier_uri'], 'String');
      }
      if (data.hasOwnProperty('subject_type')) {
        obj['subject_type'] = ApiClient.convertToType(data['subject_type'], 'String');
      }
      if (data.hasOwnProperty('token_endpoint_auth_method')) {
        obj['token_endpoint_auth_method'] = ApiClient.convertToType(data['token_endpoint_auth_method'], 'String');
      }
      if (data.hasOwnProperty('tos_uri')) {
        obj['tos_uri'] = ApiClient.convertToType(data['tos_uri'], 'String');
      }
      if (data.hasOwnProperty('userinfo_signed_response_alg')) {
        obj['userinfo_signed_response_alg'] = ApiClient.convertToType(data['userinfo_signed_response_alg'], 'String');
      }
    }
    return obj;
  }

  /**
   * AllowedCORSOrigins are one or more URLs (scheme://host[:port]) which are allowed to make CORS requests to the /oauth/token endpoint. If this array is empty, the sever's CORS origin configuration (`CORS_ALLOWED_ORIGINS`) will be used instead. If this array is set, the allowed origins are appended to the server's CORS origin configuration. Be aware that environment variable `CORS_ENABLED` MUST be set to `true` for this to work.
   * @member {Array.<String>} allowed_cors_origins
   */
  exports.prototype['allowed_cors_origins'] = undefined;
  /**
   * Audience is a whitelist defining the audiences this client is allowed to request tokens for. An audience limits the applicability of an OAuth 2.0 Access Token to, for example, certain API endpoints. The value is a list of URLs. URLs MUST NOT contain whitespaces.
   * @member {Array.<String>} audience
   */
  exports.prototype['audience'] = undefined;
  /**
   * ClientID  is the id for this client.
   * @member {String} client_id
   */
  exports.prototype['client_id'] = undefined;
  /**
   * Name is the human-readable string name of the client to be presented to the end-user during authorization.
   * @member {String} client_name
   */
  exports.prototype['client_name'] = undefined;
  /**
   * Secret is the client's secret. The secret will be included in the create request as cleartext, and then never again. The secret is stored using BCrypt so it is impossible to recover it. Tell your users that they need to write the secret down as it will not be made available again.
   * @member {String} client_secret
   */
  exports.prototype['client_secret'] = undefined;
  /**
   * SecretExpiresAt is an integer holding the time at which the client secret will expire or 0 if it will not expire. The time is represented as the number of seconds from 1970-01-01T00:00:00Z as measured in UTC until the date/time of expiration.
   * @member {Number} client_secret_expires_at
   */
  exports.prototype['client_secret_expires_at'] = undefined;
  /**
   * ClientURI is an URL string of a web page providing information about the client. If present, the server SHOULD display this URL to the end-user in a clickable fashion.
   * @member {String} client_uri
   */
  exports.prototype['client_uri'] = undefined;
  /**
   * Contacts is a array of strings representing ways to contact people responsible for this client, typically email addresses.
   * @member {Array.<String>} contacts
   */
  exports.prototype['contacts'] = undefined;
  /**
   * GrantTypes is an array of grant types the client is allowed to use.
   * @member {Array.<String>} grant_types
   */
  exports.prototype['grant_types'] = undefined;
  /**
   * @member {module:model/JSONWebKeySet} jwks
   */
  exports.prototype['jwks'] = undefined;
  /**
   * URL for the Client's JSON Web Key Set [JWK] document. If the Client signs requests to the Server, it contains the signing key(s) the Server uses to validate signatures from the Client. The JWK Set MAY also contain the Client's encryption keys(s), which are used by the Server to encrypt responses to the Client. When both signing and encryption keys are made available, a use (Key Use) parameter value is REQUIRED for all keys in the referenced JWK Set to indicate each key's intended usage. Although some algorithms allow the same key to be used for both signatures and encryption, doing so is NOT RECOMMENDED, as it is less secure. The JWK x5c parameter MAY be used to provide X.509 representations of keys provided. When used, the bare key values MUST still be present and MUST match those in the certificate.
   * @member {String} jwks_uri
   */
  exports.prototype['jwks_uri'] = undefined;
  /**
   * LogoURI is an URL string that references a logo for the client.
   * @member {String} logo_uri
   */
  exports.prototype['logo_uri'] = undefined;
  /**
   * Owner is a string identifying the owner of the OAuth 2.0 Client.
   * @member {String} owner
   */
  exports.prototype['owner'] = undefined;
  /**
   * PolicyURI is a URL string that points to a human-readable privacy policy document that describes how the deployment organization collects, uses, retains, and discloses personal data.
   * @member {String} policy_uri
   */
  exports.prototype['policy_uri'] = undefined;
  /**
   * RedirectURIs is an array of allowed redirect urls for the client, for example http://mydomain/oauth/callback .
   * @member {Array.<String>} redirect_uris
   */
  exports.prototype['redirect_uris'] = undefined;
  /**
   * JWS [JWS] alg algorithm [JWA] that MUST be used for signing Request Objects sent to the OP. All Request Objects from this Client MUST be rejected, if not signed with this algorithm.
   * @member {String} request_object_signing_alg
   */
  exports.prototype['request_object_signing_alg'] = undefined;
  /**
   * Array of request_uri values that are pre-registered by the RP for use at the OP. Servers MAY cache the contents of the files referenced by these URIs and not retrieve them at the time they are used in a request. OPs can require that request_uri values used be pre-registered with the require_request_uri_registration discovery parameter.
   * @member {Array.<String>} request_uris
   */
  exports.prototype['request_uris'] = undefined;
  /**
   * ResponseTypes is an array of the OAuth 2.0 response type strings that the client can use at the authorization endpoint.
   * @member {Array.<String>} response_types
   */
  exports.prototype['response_types'] = undefined;
  /**
   * Scope is a string containing a space-separated list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749]) that the client can use when requesting access tokens.
   * @member {String} scope
   */
  exports.prototype['scope'] = undefined;
  /**
   * URL using the https scheme to be used in calculating Pseudonymous Identifiers by the OP. The URL references a file with a single JSON array of redirect_uri values.
   * @member {String} sector_identifier_uri
   */
  exports.prototype['sector_identifier_uri'] = undefined;
  /**
   * SubjectType requested for responses to this Client. The subject_types_supported Discovery parameter contains a list of the supported subject_type values for this server. Valid types include `pairwise` and `public`.
   * @member {String} subject_type
   */
  exports.prototype['subject_type'] = undefined;
  /**
   * Requested Client Authentication method for the Token Endpoint. The options are client_secret_post, client_secret_basic, private_key_jwt, and none.
   * @member {String} token_endpoint_auth_method
   */
  exports.prototype['token_endpoint_auth_method'] = undefined;
  /**
   * TermsOfServiceURI is a URL string that points to a human-readable terms of service document for the client that describes a contractual relationship between the end-user and the client that the end-user accepts when authorizing the client.
   * @member {String} tos_uri
   */
  exports.prototype['tos_uri'] = undefined;
  /**
   * JWS alg algorithm [JWA] REQUIRED for signing UserInfo Responses. If this is specified, the response will be JWT [JWT] serialized, and signed using JWS. The default, if omitted, is for the UserInfo Response to return the Claims as a UTF-8 encoded JSON object using the application/json content-type.
   * @member {String} userinfo_signed_response_alg
   */
  exports.prototype['userinfo_signed_response_alg'] = undefined;



  return exports;
}));


