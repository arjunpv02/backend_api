import re
from rest_framework import serializers

class StudentSerializer(serializers.Serializer):
    id = serializers.IntegerField(read_only=True)
    name = serializers.CharField(max_length=100)
    age = serializers.IntegerField(min_value=3)
    email = serializers.EmailField()

    # Custom Validation for the Name field to ensure it contains only letters and spaces
    def validate_name(self, value):
        if not re.match(r'^[A-Za-z\s]+$', value):
            raise serializers.ValidationError("Name should only contain letters and spaces.")
        return value
    
    # Custom Validation for the Email field to ensure uniqueness (Example: check if email already exists)
    def validate_email(self, value):
        # Simulate a check for email uniqueness (this would typically be a database query)
        existing_emails = ['existing@example.com', 'test@example.com']  # Example of existing emails
        if value in existing_emails:
            raise serializers.ValidationError("This email is already in use.")
        return value

