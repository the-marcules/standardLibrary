package de.marcules.demo.model;

public record SignResponse(String signedPayload, String certificate, String payload) {}
