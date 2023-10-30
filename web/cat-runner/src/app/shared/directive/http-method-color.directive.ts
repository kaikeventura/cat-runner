import { Directive, ElementRef, Input, OnInit, Renderer2 } from '@angular/core';

@Directive({
  selector: '[appHttpMethodColor]'
})
export class HttpMethodColorDirective implements OnInit {
  @Input() appHttpMethodColor!: string;

  constructor(private el: ElementRef, private renderer: Renderer2) {}

  ngOnInit() {
    const httpMethod = this.appHttpMethodColor.toLowerCase();

    if (httpMethod === 'get') {
      this.renderer.setStyle(this.el.nativeElement, 'color', 'green');
    } else if (httpMethod === 'post') {
      this.renderer.setStyle(this.el.nativeElement, 'color', 'blue');
    } else if (httpMethod === 'put') {
      this.renderer.setStyle(this.el.nativeElement, 'color', 'orange');
    } else if (httpMethod === 'patch') {
      this.renderer.setStyle(this.el.nativeElement, 'color', 'purple');
    } else if (httpMethod === 'delete') {
      this.renderer.setStyle(this.el.nativeElement, 'color', 'red');
    }
  }
}