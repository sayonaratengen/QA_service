-- +goose Up
CREATE INDEX idx_answers_question_id ON answers(question_id);

-- +goose Down
DROP INDEX IF EXISTS idx_answers_question_id;
