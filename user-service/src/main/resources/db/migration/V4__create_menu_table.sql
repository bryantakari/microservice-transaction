-- Table
CREATE TABLE menu (
      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
      name         varchar(50)  NOT NULL,
      description            varchar(50)  NOT NULL,
      created_at       timestamptz  NOT NULL DEFAULT now(),
      updated_at       timestamptz  NOT NULL DEFAULT now(),
      last_login_at    timestamptz
);

CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    roles_name varchar(50) NOT NULL,
    description varchar(50) NOT NULL,
    created_at       timestamptz  NOT NULL DEFAULT now(),
    updated_at       timestamptz  NOT NULL DEFAULT now()
);

CREATE TABLE menu_mapping_role (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    menu_id UUID NOT NULL,
    roles_id UUID NOT NULL,

    CONSTRAINT fk_menu_mapping_role
    FOREIGN KEY (menu_id)
    REFERENCES menu(id)
);

CREATE TABLE accounts_mapping_role (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role_id UUID NOT NULL,
    account_id UUID NOT NULL,

    CONSTRAINT fk_accounts_mapping_role
    FOREIGN KEY (account_id)
    REFERENCES accounts(id)
);