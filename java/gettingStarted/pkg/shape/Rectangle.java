package pkg.shape;

public class Rectangle implements ShapeInterface {
     double a;
     double b;

     public Rectangle(double a, double b) {
        this.a = a;
        this.b = b;
     }
    
     public Rectangle(double a) {
        this.a = a;
        this.b = a;
     }
    
    public double getA() {
        return a;
    }

    public double getB() {
        return b;
    }

    @Override
    public double area() {
       return this.a*this.b;
    }

    @Override
    public double circumfence() {
        return 2*(this.a + this.b);
    }
    
}
