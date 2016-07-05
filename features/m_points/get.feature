@m_point
Feature: ポイント情報を取得

  @success
  Scenario Outline: 存在するアカウントにログイン
    When I send and accept JSON
    And I send a GET request to "m_points"
    Then the response status should be "200"
    And the JSON response should follow "features/schemas/success_login.json"

  @failure
  Scenario Outline: Check unsupported methods
    When I send and accept JSON
    And I send a <method> request to "m_points"
    Then the response status should be "405"
    Examples:
      | method |
      | POST   |
      | PUT    |
      | DELETE |
