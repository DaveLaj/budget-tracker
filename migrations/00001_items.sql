-- +goose Up
-- +goose StatementBegin
CREATE TABLE `items` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `description` varchar(255) NOT NULL,
    `amount` int(11) NOT NULL,
    `date` date NOT NULL,

    PRIMARY KEY (`id`)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `items`;
-- +goose StatementEnd
