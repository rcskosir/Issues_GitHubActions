`query=
					mutation (
                      $project:ID!, $item:ID!,
                      $status_field:ID!, $status_value:String!,
                      $pr_field:ID!, $pr_value:String!,
                      $user_field:ID!, $user_value:String!,
                      $daysOpen_field:ID!, $daysOpen_value:Float!,
                      $daysWait_field:ID!, $daysWait_value:Float!,
					) {
					  set_status: updateProjectV2ItemFieldValue(input: {
						projectId: $project
						itemId: $item
						fieldId: $status_field
						value: {
						  singleSelectOptionId: $status_value
						  }
					  }) {
						projectV2Item {
						  id
						  }
					  }
					  set_pr: updateProjectV2ItemFieldValue(input: {
						projectId: $project
						itemId: $item
						fieldId: $pr_field
						value: {
						  text: $pr_value
						}
					  }) {
						projectV2Item {
						  id
						  }
					  }
                      set_user: updateProjectV2ItemFieldValue(input: {
						projectId: $project
						itemId: $item
						fieldId: $user_field
						value: {
						  text: $user_value
						}
					  }) {
						projectV2Item {
						  id
						  }
					  }
					  set_dopen: updateProjectV2ItemFieldValue(input: {
						projectId: $project
						itemId: $item
						fieldId: $daysOpen_field
						value: {
						  number: $daysOpen_value
						}
					  }) {
						projectV2Item {
						  id
						  }
					  }
					  set_dwait: updateProjectV2ItemFieldValue(input: {
						projectId: $project
						itemId: $item
						fieldId: $daysWait_field
						value: {
						  number: $daysWait_value
						}
					  }) {
						projectV2Item {
						  id
						  }
					  }
					}
				`
variables='{
  "eventId": "276754274"
}'
query="$(echo $query)"
curl -X POST https://api.meetup.com/gql \
  -H 'Authorization: Bearer {YOUR_TOKEN}' \
  -H 'Content-Type: application/json' \
  -d @- <<EOF
      {"query": "$query", "variables": "$variables"}