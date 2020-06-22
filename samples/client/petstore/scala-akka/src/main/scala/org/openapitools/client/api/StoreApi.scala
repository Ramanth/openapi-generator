/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
package org.openapitools.client.api

import org.openapitools.client.model.Order
import org.openapitools.client.core._
import org.openapitools.client.core.CollectionFormats._
import org.openapitools.client.core.ApiKeyLocations._

object StoreApi {

  def apply(baseUrl: String = "http://petstore.swagger.io/v2") = new StoreApi(baseUrl)
}

class StoreApi(baseUrl: String) {
  
  /**
   * For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors
   * 
   * Expected answers:
   *   code 400 :  (Invalid ID supplied)
   *   code 404 :  (Order not found)
   * 
   * @param orderId ID of the order that needs to be deleted
   */
  def deleteOrder(orderId: String): ApiRequest[Unit] =
    ApiRequest[Unit](ApiMethods.DELETE, baseUrl, "/store/order/{orderId}", "application/json")
      .withPathParam("orderId", orderId)
      .withErrorResponse[Unit](400)
      .withErrorResponse[Unit](404)
      

  /**
   * Returns a map of status codes to quantities
   * 
   * Expected answers:
   *   code 200 : Map[String, Int] (successful operation)
   * 
   * Available security schemes:
   *   api_key (apiKey)
   */
  def getInventory()(implicit apiKey: ApiKeyValue): ApiRequest[Map[String, Int]] =
    ApiRequest[Map[String, Int]](ApiMethods.GET, baseUrl, "/store/inventory", "application/json")
      .withApiKey(apiKey, "api_key", HEADER)
      .withSuccessResponse[Map[String, Int]](200)
      

  /**
   * For valid response try integer IDs with value <= 5 or > 10. Other values will generated exceptions
   * 
   * Expected answers:
   *   code 200 : Order (successful operation)
   *   code 400 :  (Invalid ID supplied)
   *   code 404 :  (Order not found)
   * 
   * @param orderId ID of pet that needs to be fetched
   */
  def getOrderById(orderId: Long): ApiRequest[Order] =
    ApiRequest[Order](ApiMethods.GET, baseUrl, "/store/order/{orderId}", "application/json")
      .withPathParam("orderId", orderId)
      .withSuccessResponse[Order](200)
      .withErrorResponse[Unit](400)
      .withErrorResponse[Unit](404)
      

  /**
   * Expected answers:
   *   code 200 : Order (successful operation)
   *   code 400 :  (Invalid Order)
   * 
   * @param order order placed for purchasing the pet
   */
  def placeOrder(order: Order): ApiRequest[Order] =
    ApiRequest[Order](ApiMethods.POST, baseUrl, "/store/order", "application/json")
      .withBody(order)
      .withSuccessResponse[Order](200)
      .withErrorResponse[Unit](400)
      



}

