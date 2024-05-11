Feature: Migrator
  In order to migrate the database
  As a user
  I want to be able to migrate the database

  Scenario: Migrate the databases
    Given I have to migrate the databases
    When I migrate the databases
    Then the databases should be migrated
