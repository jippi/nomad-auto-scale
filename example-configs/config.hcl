// Nomad Client configuration
nomad {
  address = "http://localhost:4646"
}

backend "rabbitmq" {
  type     = "rabbitmq"
  address  = "http://localhost:15672"
  username = "golang-demo"
  password = "golang-demo"
}

// Scale for the job "content-themes"
job "content-themes" {
  // For group "populator"
  group "populator" {
    // (required) The minimum nuber of tasks to run for this job
    min_count = 0

    // (required) The maximum number of tasks to run for this job
    max_count = 1

    // Scale by a rule
    // Again the name does not matter for execution
    rule "rabbitmq queue size" {
      // (required) What backend to use, this will define which configuration
      // is valid and which checks you can execute
      backend = "rabbitmq"

      // (required) The check type, could be anything implemented by the backend
      //
      // Example for rabbitmq
      //   - "queue_length"
      //   - "queue_utilization"
      //   - "consumer_count"
      check_type = "queue_length"

      // (required) The comparison do do, this supports the basic match operations like
      //
      //   - above (>)
      //   - below (<)
      //   - equal (==)
      //   - not_equal (!=)
      //   - above_or_equal (>=)
      //   - below_or_equal (<=)
      comparison = "above"

      // (required) The value to compare to, this should be a float or integer
      comparison_value = 0.9

      // (optional) You can define how often this rule should be checked, by default it will me checked every minte
      cron = "*/15 * * *"

      // (optional) Possible actions to take if the comparison is evaluated to "true"
      if_true {
        // Boolean flag to decide if other rules should be processed or not
        final        = false
        action       = "increase_count"
        action_value = 0

        // A sample action to send notifcation to slack
        notify {
          type    = "slack"
          room    = "#developers"
          message = "${SCALE_NAME} moved from ${PREVIOUS_COUNT} to ${CURRENT_COUNT}"
        }
      }

      // (optional) Possible actions to take if the comparison is evaluated to "false"
      // the keys will match exactly what if_true can do
      if_false {}
    }
  }
}
