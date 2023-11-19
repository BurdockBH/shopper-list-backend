-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS shopping_list_table (
                                                 `id` INT NOT NULL AUTO_INCREMENT,
                                                 `shopper_id` INT NULL,
                                                 `item_id` INT NULL,
                                                 PRIMARY KEY (`id`),
                                                 INDEX `shopper_id_idx` (`shopper_id` ASC) VISIBLE,
                                                 INDEX `item_id_idx` (`item_id` ASC) VISIBLE,
                                                 CONSTRAINT `shopper_id`
                                                     FOREIGN KEY (`shopper_id`)
                                                         REFERENCES `shopping-list`.`shopper` (`id`)
                                                         ON DELETE NO ACTION
                                                         ON UPDATE NO ACTION,
                                                 CONSTRAINT `item_id`
                                                     FOREIGN KEY (`item_id`)
                                                         REFERENCES `shopping-list`.`items` (`id`)
                                                         ON DELETE NO ACTION
                                                         ON UPDATE NO ACTION);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shopping_list_table;
-- +goose StatementEnd
