@u_point
Feature: 各ユーザーのポイントを取得
  @success
  Scenario: ログインユーザのすべてのポイント情報を取得
    When I am logged in as:
      | token | valid_token |
      | app_code | gcpn |
    And I send a GET request to "accounts/{account_id}/points"
    Then the response status should be "200"
    And the JSON response should follow "features/schemas/get_u_points.json"

  @authorize
  Scenario: ログインユーザではないユーザのポイント情報を取得
    When I am logged in as:
      | token    | sms_valid_token |
      | app_code | gcpn        |
    And I send a GET request to "accounts/1/points"
    Then the response status should be "403"

  @authorize
  Scenario: ログインせずにユーザーのポイント情報を取得
    When I send and accept JSON
    And I send a GET request to "accounts/7/points"
    Then the response status should be "403"
