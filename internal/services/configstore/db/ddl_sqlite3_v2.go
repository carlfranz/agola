// Code generated by go generate; DO NOT EDIT.
package db

var DDLSqlite3V2 = []string{
	"create table if not exists remotesource (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, name varchar NOT NULL, apiurl varchar NOT NULL, skip_verify integer NOT NULL, type varchar NOT NULL, auth_type varchar NOT NULL, oauth2_client_id varchar NOT NULL, oauth2_client_secret varchar NOT NULL, ssh_host_key varchar NOT NULL, skip_ssh_host_key_check integer NOT NULL, registration_enabled integer NOT NULL, login_enabled integer NOT NULL, PRIMARY KEY (id))",
	"create table if not exists user_t (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, name varchar NOT NULL, secret varchar NOT NULL, admin integer NOT NULL, PRIMARY KEY (id))",
	"create table if not exists usertoken (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, user_id varchar NOT NULL, name varchar NOT NULL, value varchar NOT NULL, PRIMARY KEY (id), foreign key (user_id) references user_t(id))",
	"create table if not exists linkedaccount (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, user_id varchar NOT NULL, remote_user_id varchar NOT NULL, remote_user_name varchar NOT NULL, remote_user_avatar_url varchar NOT NULL, remote_source_id varchar NOT NULL, user_access_token varchar NOT NULL, oauth2_access_token varchar NOT NULL, oauth2_refresh_token varchar NOT NULL, oauth2_access_token_expires_at timestamp NOT NULL, PRIMARY KEY (id), foreign key (user_id) references user_t(id), foreign key (remote_source_id) references remotesource(id))",
	"create table if not exists organization (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, name varchar NOT NULL, visibility varchar NOT NULL, creator_user_id varchar NOT NULL, PRIMARY KEY (id))",
	"create table if not exists orgmember (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, organization_id varchar NOT NULL, user_id varchar NOT NULL, member_role varchar NOT NULL, PRIMARY KEY (id), foreign key (organization_id) references organization(id), foreign key (user_id) references user_t(id))",
	"create table if not exists projectgroup (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, name varchar NOT NULL, parent_kind varchar NOT NULL, parent_id varchar NOT NULL, visibility varchar NOT NULL, PRIMARY KEY (id))",
	"create table if not exists project (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, name varchar NOT NULL, parent_kind varchar NOT NULL, parent_id varchar NOT NULL, secret varchar NOT NULL, visibility varchar NOT NULL, remote_repository_config_type varchar NOT NULL, remote_source_id varchar NOT NULL, linked_account_id varchar NOT NULL, repository_id varchar NOT NULL, repository_path varchar NOT NULL, ssh_private_key varchar NOT NULL, skip_ssh_host_key_check integer NOT NULL, webhook_secret varchar NOT NULL, pass_vars_to_forked_pr integer NOT NULL, default_branch varchar NOT NULL, PRIMARY KEY (id))",
	"create table if not exists secret (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, name varchar NOT NULL, parent_kind varchar NOT NULL, parent_id varchar NOT NULL, type varchar NOT NULL, data text NOT NULL, secret_provider_id varchar NOT NULL, path varchar NOT NULL, PRIMARY KEY (id))",
	"create table if not exists variable (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, name varchar NOT NULL, parent_kind varchar NOT NULL, parent_id varchar NOT NULL, variable_values text NOT NULL, PRIMARY KEY (id))",
	"create table if not exists orginvitation (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, user_id varchar NOT NULL, organization_id varchar NOT NULL, role varchar NOT NULL, PRIMARY KEY (id), foreign key (user_id) references user_t(id), foreign key (organization_id) references organization(id))",

	// indexes
}
