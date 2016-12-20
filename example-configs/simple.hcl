scale "sendgrid.batch_sync" {
  min_count = 0
  max_count = 1

  rule "rabbitmq queue size" {
    backend          = "rabbitmq"
    check_type       = "queue_length"
    queue_name       = "sendgrid.batch"
    comparison       = "above"
    comparison_value = 0.9

    if_true {
      increase_count_by = 1
    }

    if_false {
      decrase_count_by = 1
    }
  }
}
