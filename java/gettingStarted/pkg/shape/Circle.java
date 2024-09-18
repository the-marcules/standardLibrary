package pkg.shape;

public class Circle implements ShapeInterface {
    double radius;

    public Circle(double r) {
        this.radius = r;
    }

    public double getRadius() {
        return radius;
    }

    @Override
    public double area() {
        return Math.PI * Math.pow(radius, 2);
    }

    @Override
    public double circumfence() {
        return Math.PI * this.radius * 2;
    }
    
}
