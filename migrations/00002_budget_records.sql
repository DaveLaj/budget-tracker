-- +goose Up
-- +goose StatementBegin
CREATE TABLE `budget_records` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` int(11) NOT NULL,
    `DateTimeIn` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `date` date NOT NULL,
    `amount` int(11) NOT NULL,
    `item_id` int(11) NOT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`item_id`) REFERENCES `items`(`id`) 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `budget_records`;
-- +goose StatementEnd
