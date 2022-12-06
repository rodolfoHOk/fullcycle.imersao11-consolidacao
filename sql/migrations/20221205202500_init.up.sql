CREATE TABLE matches (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  match_date DATETIME,
  team_a_id VARCHAR(36),
  team_a_name VARCHAR(255),
  team_b_id VARCHAR(36),
  team_b_name VARCHAR(255),
  result VARCHAR(255)
);

CREATE TABLE players (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price DECIMAL(10,2) NOT NULL
);

CREATE TABLE teams (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE team_players (
  team_id VARCHAR(36) NOT NULL,
  player_id VARCHAR(36) NOT NULL
);

CREATE TABLE my_team (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  score DECIMAL(10,2) NOT NULL
);

CREATE TABLE my_team_players (
  my_team_id VARCHAR(36) NOT NULL,
  player_id VARCHAR(36) NOT NULL
);

CREATE TABLE actions (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  match_id VARCHAR(36) NOT NULL,
  team_id VARCHAR(36) NOT NULL,
  player_id VARCHAR(36) NOT NULL,
  action VARCHAR(255) NOT NULL,
  minute INTEGER NOT NULL,
  score DECIMAL(10,2) NOT NULL
);
