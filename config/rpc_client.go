package config

type RPCClient struct {
	UserServiceName     string `mapstructure:"user_service_name" yaml:"user_service_name"`
	PublishServiceName  string `mapstructure:"publish_service_name" yaml:"publish_service_name"`
	FeedServiceName     string `mapstructure:"feed_service_name" yaml:"feed_service_name"`
	FavoriteServiceName string `mapstructure:"favorite_service_name" yaml:"favorite_service_name"`
	CommentServiceName  string `mapstructure:"comment_service_name" yaml:"comment_service_name"`
	RelationServiceName string `mapstructure:"relation_service_name" yaml:"relation_service_name"`
	MessageServiceName  string `mapstructure:"message_service_name" yaml:"message_service_name"`
}
