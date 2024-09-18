import boto3

# CloudWatch-Client erstellen
client = boto3.client('cloudwatch')

# Alarme auflisten
response = client.describe_alarms()

# Alarmliste ausgeben
for alarm in response['MetricAlarms']:
    state = "🚨"
    if alarm['StateValue'] == "OK":
        state = "🟢"
    print(state + " --> " + alarm['AlarmName'])
