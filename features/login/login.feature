@login
Feature: ログイントークンでのログイン

  @success
  Scenario Outline: 存在するアカウントにログイン
    When I send and accept JSON
    And I set form request body to:
    | login_token | <token>    |
    | app_code    | <app_code> |
    And I send a POST request to "login"
    Then the response status should be "200"
    And the JSON response should follow "features/schemas/success_login.json"
    Examples:
      | token       | app_code |
      | valid_token | gcpn     |

  @failure
  Scenario Outline: 存在しないパラメータでログイン
    When I send and accept JSON
    And I set form request body to:
    | login_token | <token>    |
    | app_code    | <app_code> |
    And I send a POST request to "login"
    Then the response status should be "400"
    Examples:
      | token        | app_code |
      | invalidtoken | gcpn     |

  @failure
  Scenario Outline: Check unsupported methods
    When I send and accept JSON
    And I send a <method> request to "login"
    Then the response status should be "405"
    Examples:
      | method |
      | GET    |
      | PUT    |
      | DELETE |
